package cron

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joaosoft/dbr"
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
)

type Cron struct {
	jobs       MapJobs
	schedulers MapSchedule
	location   *time.Location

	dbr           *dbr.Dbr
	config        *CronConfig
	isLogExternal bool
	pm            *manager.Manager
	logger        logger.ILogger
	quit          chan int
	mux           *sync.Mutex
}

// Start ...
func New(options ...CronOption) (*Cron, error) {
	quit := make(chan int)
	config, simpleConfig, err := NewConfig()

	cron := &Cron{
		quit:       quit,
		pm:         manager.NewManager(manager.WithRunInBackground(true), manager.WithQuitChannel(quit)),
		logger:     logger.NewLogDefault("cron", logger.WarnLevel),
		config:     config.Cron,
		jobs:       make(MapJobs),
		schedulers: make(MapSchedule),
		location:   time.Local,
		mux:        &sync.Mutex{},
	}

	if cron.isLogExternal {
		cron.pm.Reconfigure(manager.WithLogger(cron.logger))
	}

	if err != nil {
		cron.logger.Error(err.Error())
	} else if config.Cron != nil {
		cron.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Cron.Log.Level)
		cron.logger.Debugf("setting log level to %s", level)
		cron.logger.Reconfigure(logger.WithLevel(level))
	}

	cron.Reconfigure(options...)

	// execute migrations
	if cron.config.Migration != nil {
		migrationService, err := migration.NewCmdService(migration.WithCmdConfiguration(cron.config.Migration))
		if err != nil {
			return nil, err
		}

		if _, err := migrationService.Execute(migration.OptionUp, 0, migration.ExecutorModeDatabase); err != nil {
			return nil, err
		}
	}

	cron.dbr, err = dbr.New(dbr.WithConfiguration(cron.config.Dbr))
	if err != nil {
		return nil, err
	}

	// load jobs
	if err = cron.loadJobs(); err != nil {
		return nil, err
	}

	// schedule jobs
	if err = cron.scheduleJobs(); err != nil {
		return nil, err
	}

	// start system
	if err = cron.pm.Start(); err != nil {
		return nil, err
	}

	cron.logger.Info("cron service started!")
	return cron, nil
}

func (cron *Cron) Wait() (err error) {
	cron.logger.Info("waiting for quit signal...")
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	<-termChan
	cron.logger.Info("quitting cron")
	cron.quit <- 1

	return nil
}

func (cron *Cron) AddJob(job *Job) *Cron {
	cron.mux.Lock()
	defer cron.mux.Unlock()

	cron.jobs[job.Key] = job

	if err := cron.scheduleJob(job); err != nil {
		cron.logger.Errorf("error scheduling job with name: %s", job.Name)
	}

	return cron
}

func (cron *Cron) AddJobTask(key string, task ITask) *Cron {
	cron.mux.Lock()
	defer cron.mux.Unlock()

	if job, ok := cron.jobs[key]; ok {
		job.Tasks = append(job.Tasks, task)
	}
	return cron
}

func (cron *Cron) AddJobTaskWithFuncs(key string, canExecuteFunc CheckExecuteFunc, beforeFunc ExecuteFunc, middleFunc ExecuteFunc, afterFunc ExecuteFunc) *Cron {
	cron.mux.Lock()
	defer cron.mux.Unlock()

	if job, ok := cron.jobs[key]; ok {
		job.Tasks = append(job.Tasks, NewTaskGeneric(canExecuteFunc, beforeFunc, middleFunc, afterFunc))
	}
	return cron
}

func (cron *Cron) loadJobs() error {
	var jobs ListJobs

	_, err := cron.dbr.Select("*").
		From(cronTableJob).
		Where("active").
		OrderAsc("position").
		Load(&jobs)

	for _, job := range jobs {
		cron.jobs[job.Key] = job
	}

	return err
}

func (cron *Cron) scheduleJobs() (err error) {
	for _, job := range cron.jobs {
		if err = cron.scheduleJob(job); err != nil {
			return err
		}
	}

	return nil
}

func (cron *Cron) scheduleJob(job *Job) (err error) {
	cron.logger.Infof("scheduling job with name: %s", job.Name)
	schedule := cron.NewSchedule(job)
	cron.schedulers[job.Key] = schedule

	if err = cron.pm.AddProcess(job.Key, schedule); err != nil {
		return err
	}

	return nil
}
