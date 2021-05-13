package cron

import "errors"

var (
	ErrorInvalidTask = errors.New("invalid function")
	ErrorInvalidFunction                = errors.New("invalid task")
)
