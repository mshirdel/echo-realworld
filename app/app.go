package app

import (
	"fmt"

	"github.com/mshirdel/echo-realworld/app/db"
	"github.com/mshirdel/echo-realworld/config"
	"github.com/sirupsen/logrus"
)

type Application struct {
	configPath string
	Cfg        *config.Config
	DB         *db.DB
}

func New(configPath string) *Application {
	return &Application{
		configPath: configPath,
	}
}

func (a *Application) InitAll() error {
	if err := a.initConfig(); err != nil {
		return fmt.Errorf("error in initializing config: %w", err)
	}

	if err := a.initDatabase(); err != nil {
		return fmt.Errorf("error in initializing database: %w", err)
	}

	return nil
}

func (a *Application) Shutdown() {
	logrus.Info("shutting down application....")
	// close things like DB, tracing and etc
	a.DB.Close()
}

func (a *Application) initConfig() (err error) {
	if a.Cfg != nil {
		return nil
	}

	a.Cfg, err = config.Init(a.configPath)
	if err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}

	return
}

func (a *Application) initDatabase() error {
	if a.DB != nil {
		return nil
	}

	a.DB = db.New(a.Cfg)
	if err := a.DB.InitDB(); err != nil {
		return fmt.Errorf("error in initializing db: %w", err)
	}

	return nil
}
