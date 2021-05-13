package cron

type TaskGeneric struct {
	canExecuteFunc CheckExecuteFunc
	funcs     []ExecuteFunc
}

func NewTaskGeneric(canExecuteFunc CheckExecuteFunc, f ExecuteFunc, fs ...ExecuteFunc) *TaskGeneric {
	return &TaskGeneric{
		canExecuteFunc: canExecuteFunc,
		funcs:     append([]ExecuteFunc{f}, fs...),
	}
}

func (task *TaskGeneric) CanExecute() (bool, error) {
	if task.canExecuteFunc == nil {
		return true, nil
	}
	return task.canExecuteFunc()
}

func (task *TaskGeneric) Execute(breakOnError bool) (err error) {
	if task == nil {
		return ErrorInvalidTask
	}

	for _, f := range task.funcs {
		if f == nil {
			return ErrorInvalidFunction
		}
		if err = f(); err != nil && breakOnError {
			return err
		}
	}

	return nil
}
