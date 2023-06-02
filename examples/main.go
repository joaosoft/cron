package main

import (
	"github.com/joaosoft/cron"
	"github.com/joaosoft/logger"
)

func main() {
	// create cron
	myCron, err := cron.New()
	if err != nil {
		panic(err)
	}

	// add tasks to jobs
	myCron.
		AddJobTaskWithFuncs("teste1_key", canExecute, a, b, c).
		AddJobTask("teste2_key", cron.NewTaskGeneric(nil, b, nil, c)).
		AddJobTaskWithFuncs("teste3_key", canExecute, a, b, c)

	// wait
	myCron.Wait()
}

func canExecute() (bool, error) {
	logger.Info("executing can execute method")
	return true, nil
}

func a() error {
	logger.Info("executing a method")
	return nil
}

func b() error {
	logger.Info("executing b method")
	return nil
}

func c() error {
	logger.Info("executing c method")
	return nil
}
