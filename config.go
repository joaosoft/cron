package cron

import (
	"fmt"

	"github.com/joaosoft/dbr"

	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
)

// AppConfig ...
type AppConfig struct {
	Cron *CronConfig `json:"cron"`
}

// CronConfig ...
type CronConfig struct {
	Migration *migration.MigrationConfig `json:"migration"`
	Dbr       *dbr.DbrConfig             `json:"dbr"`
	Log       struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}
