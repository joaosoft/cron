package cron

type ListTasks []ITask

type CheckExecuteFunc func() (bool, error)
type ExecuteFunc func() error

type ITask interface {
	CanExecute() (bool, error)
	Before() error
	Middle() error
	After() error
}

func (tasks ListTasks) Execute() (err error) {
	for _, task := range tasks {
		canExecute, err := task.CanExecute()
		if err != nil {
			return err
		}

		if !canExecute {
			return nil
		}

		if err = task.Before(); err != nil {
			return err
		}

		if err = task.Middle(); err != nil {
			return err
		}

		if err = task.After(); err != nil {
			return err
		}
	}
	return nil
}
