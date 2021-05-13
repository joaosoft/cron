package cron

type ListTasks []ITask

type CheckExecuteFunc func() (bool, error)
type ExecuteFunc func() error

type ITask interface {
	CanExecute() (bool, error)
	Execute(breakOnError bool) error
}

func (tasks ListTasks) Execute() (err error) {
	var canExecute bool
	for _, task := range tasks {
		if task == nil {
			return ErrorInvalidTask
		}
		canExecute, err = task.CanExecute()
		if err != nil {
			return err
		}

		if !canExecute {
			return nil
		}

		if err = task.Execute(false); err != nil {
			return err
		}
	}
	return nil
}
