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
	return task.canExecuteFunc()
}

func (task *TaskGeneric) Before() error {
	return task.beforeFunc()
}

func (task *TaskGeneric) Middle() error {
	return task.middleFunc()
}

func (task *TaskGeneric) After() error {
	return task.afterFunc()
}
