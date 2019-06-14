package cron

import (
	"sync"
	"time"
)

type MapSchedule map[string]*Schedule

type Schedule struct {
	cron    *Cron
	Job     *Job
	timer   *timer
	started bool
	quit    chan bool
}

func (cron *Cron) NewSchedule(job *Job) *Schedule {
	return &Schedule{
		quit:  make(chan bool),
		cron:  cron,
		Job:   job,
		timer: newTimer(cron.location, job.Settings),
	}
}

func (schedule *Schedule) Start(waitGroup ...*sync.WaitGroup) error {
	var wg *sync.WaitGroup

	if len(waitGroup) == 0 {
		wg = &sync.WaitGroup{}
		wg.Add(1)
	} else {
		wg = waitGroup[0]
	}

	defer wg.Done()

	var err error
	schedule.started = true
	waitTimeForNext := schedule.timer.Next()

	go func() {
		for {
			schedule.cron.logger.Infof("job %s waiting next scheduled time %s]", schedule.Job.Key, waitTimeForNext.String())

			select {
			case <-time.After(waitTimeForNext):

				schedule.cron.logger.Infof("executing job %s", schedule.Job.Key)
				if err = schedule.Job.Tasks.Execute(); err != nil {
					schedule.cron.logger.Errorf("error executing job %s [error: %s]", schedule.Job.Key, err.Error())
				}
				waitTimeForNext = schedule.timer.Next()

			case <-schedule.quit:
				schedule.cron.logger.Infof("quitting job %s", schedule.Job.Key)
				return
			}
		}
	}()

	return nil
}

func (schedule *Schedule) Stop(waitGroup ...*sync.WaitGroup) error {
	var wg *sync.WaitGroup

	if len(waitGroup) == 0 {
		wg = &sync.WaitGroup{}
		wg.Add(1)
	} else {
		wg = waitGroup[0]
	}

	defer wg.Done()

	schedule.cron.logger.Infof("stopping job %s", schedule.Job.Key)
	schedule.started = false
	schedule.quit <- true

	return nil
}

func (schedule *Schedule) Started() bool {
	return schedule.started
}
