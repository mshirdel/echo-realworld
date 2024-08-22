package app

import (
	"fmt"

	"github.com/mshirdel/echo-realworld/app/db"
	"github.com/mshirdel/echo-realworld/app/repo"
	"github.com/mshirdel/echo-realworld/app/service"
	"github.com/mshirdel/echo-realworld/config"
	"github.com/sirupsen/logrus"
)

type Application struct {
	configPath string
	Cfg        *config.Config
	Database   *db.DB
	Repo       *repo.Repository
	Svc        *service.Service
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

	a.initRepositories()
	
	if err := a.initServices(); err != nil {
		return fmt.Errorf("error in initializing services: %w", err)
	}

	return nil
}

func (a *Application) initServices() error {
	if a.Svc != nil {
		return nil
	}

	a.Svc = service.New(a.Cfg, a.Repo)
	return a.Svc.InitAll()
}

func (a *Application) Shutdown() {
	logrus.Info("shutting down application....")
	// close things like DB, tracing and etc
	a.Database.Close()
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
	if a.Database != nil {
		return nil
	}

	a.Database = db.New(a.Cfg)
	if err := a.Database.Init(); err != nil {
		return fmt.Errorf("error in initializing db: %w", err)
	}

	return nil
}

func (a *Application) initRepositories() {
	if a.Repo != nil {
		return
	}

	a.Repo = repo.New(a.Cfg, a.Database)
	a.Repo.InitAll()
}
