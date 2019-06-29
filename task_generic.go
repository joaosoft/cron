package cron

type TaskGeneric struct {
	canExecuteFunc CheckExecuteFunc
	beforeFunc     ExecuteFunc
	middleFunc     ExecuteFunc
	afterFunc      ExecuteFunc
}

func NewTaskGeneric(canExecuteFunc CheckExecuteFunc, beforeFunc ExecuteFunc, middleFunc ExecuteFunc, afterFunc ExecuteFunc) *TaskGeneric {
	return &TaskGeneric{
		canExecuteFunc: canExecuteFunc,
		beforeFunc:     beforeFunc,
		middleFunc:     middleFunc,
		afterFunc:      afterFunc,
	}
}

func (task *TaskGeneric) CanExecute() (bool, error) {
	if task.canExecuteFunc == nil {
		return true, nil
	}
	return task.canExecuteFunc()
}

func (task *TaskGeneric) Before() error {
	if task.beforeFunc == nil {
		return nil
	}
	return task.beforeFunc()
}

func (task *TaskGeneric) Middle() error {
	if task.middleFunc == nil {
		return nil
	}
	return task.middleFunc()
}

func (task *TaskGeneric) After() error {
	if task.afterFunc == nil {
		return nil
	}
	return task.afterFunc()
}
