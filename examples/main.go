package main

import (
	"cron"
	"logger"
)

func main() {
	// create cron
	myCron, err := cron.New()
	if err != nil {
		panic(err)
	}

	// add tasks to jobs
	myCron.
		AddJobTaskWithFuncs("teste1_key", canExecute, before, middle, after).
		AddJobTask("teste2_key", cron.NewTaskGeneric(canExecute, before, middle, after)).
		AddJobTaskWithFuncs("teste3_key", canExecute, before, middle, after)

	// wait
	myCron.Wait()
}

func canExecute() (bool, error) {
	logger.Info("executing can execute method")
	return true, nil
}

func before() error {
	logger.Info("executing before method")
	return nil
}

func middle() error {
	logger.Info("executing middle method")
	return nil
}

func after() error {
	logger.Info("executing after method")
	return nil
}
