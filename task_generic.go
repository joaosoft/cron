package cron

type TaskGeneric struct {
	canExecuteFunc CheckExecuteFunc
	funcs     []ExecuteFunc
}

func NewTaskGeneric(canExecuteFunc CheckExecuteFunc, f1 ExecuteFunc, fn ...ExecuteFunc) *TaskGeneric {
	return &TaskGeneric{
		canExecuteFunc: canExecuteFunc,
		funcs:     append([]ExecuteFunc{f1}, fn...),
	}
}

func (task *TaskGeneric) CanExecute() (bool, error) {
	if task.canExecuteFunc == nil {
		return true, nil
	}
	return task.canExecuteFunc()
}

func (task *TaskGeneric) Execute(breakOnError bool) (err error) {
	for _, f := range task.funcs {
		if err = f(); breakOnError && err != nil {
			return err
		}
	}

	return nil
}
