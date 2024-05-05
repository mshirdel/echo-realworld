package app

import (
	"fmt"

	"github.com/mshirdel/echo-realworld/config"
	"github.com/sirupsen/logrus"
)

type Application struct {
	configPath string
	Cfg        *config.Config
}

func New(configPath string) *Application {
	return &Application{
		configPath: configPath,
	}
}

func (a *Application) InitAll() error {
	err := a.initConfig()

	return err
}

func (a *Application) Shutdown() {
	logrus.Info("shutting down application....")
	// close things like DB, tracing and etc
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
