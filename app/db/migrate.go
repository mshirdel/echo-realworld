package db

import (
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"gorm.io/gorm"

	// read migrations from filesystem
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMigrate(gormdb *gorm.DB, dbname, path string) (*migrate.Migrate, error) {
	db, err := gormdb.DB()
	if err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{
		DatabaseName:    dbname,
		MigrationsTable: "schema_migrations",
	})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance("file://"+path, dbname, driver)
}
