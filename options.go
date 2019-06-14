package cron

import (
	"time"

	logger "github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// CronOption ...
type CronOption func(cron *Cron)

// Reconfigure ...
func (cron *Cron) Reconfigure(options ...CronOption) {
	for _, option := range options {
		option(cron)
	}
}

// WithConfiguration ...
func WithConfiguration(config *CronConfig) CronOption {
	return func(cron *Cron) {
		cron.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) CronOption {
	return func(cron *Cron) {
		cron.logger = logger
		cron.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) CronOption {
	return func(cron *Cron) {
		cron.logger.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) CronOption {
	return func(cron *Cron) {
		cron.pm = mgr
	}
}

// WithLocation ...
func WithLocation(location *time.Location) CronOption {
	return func(cron *Cron) {
		cron.location = location
	}
}
