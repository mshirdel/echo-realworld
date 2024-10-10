package cmd

import (
	"errors"
	"fmt"
	libmigrate "github.com/golang-migrate/migrate/v4"
	"github.com/mshirdel/echo-realworld/app"
	"github.com/mshirdel/echo-realworld/app/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	_migrationsPath string
	_steps          int
	_force          int
)

var _migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	RunE:  migrate,
}

func init() {
	_migrateCmd.Flags().StringVarP(&_migrationsPath, "migrations-path", "m", "./migrations", "path to migrations directory")
	_migrateCmd.Flags().IntVarP(&_steps, "steps", "s", 0, "number of steps to migrate. positive steps for up and negative steps for down. zero to upgrade all.")
	_migrateCmd.Flags().IntVar(&_force, "force", 0, "Force sets a migration version. It resets the dirty state to false.")
}

func migrate(_ *cobra.Command, _ []string) error {
	app := app.New(_cfgFile)
	defer app.Shutdown()

	if err := app.InitConfig(); err != nil {
		return err
	}

	if err := app.InitDatabase(); err != nil {
		return err
	}

	m, err := db.NewMigrate(app.Database.Realworld, app.Cfg.Database.DBName, _migrationsPath)
	if err != nil {
		return err
	}

	logrus.Info("migration started ...")

	if _force > 0 {
		if err := m.Force(_force); err != nil {
			return err
		}
	}

	if _steps == 0 {
		err = m.Up()
	} else {
		err = m.Steps(_steps)
	}

	if err != nil {
		if errors.Is(err, libmigrate.ErrNoChange) {
			logrus.Info("no migration to change")
		} else {
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	logrus.Info("migration finished.")

	return nil
}
