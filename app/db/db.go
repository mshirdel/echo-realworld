package db

import (
	"fmt"
	"strings"

	"github.com/mshirdel/echo-realworld/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	cfg       *config.Database
	Realworld *gorm.DB
}

func New(cfg *config.Config) *DB {
	return &DB{cfg: &cfg.Database}
}

func (d *DB) Init() error {
	var err error
	if d.Realworld != nil {
		return err
	}

	d.Realworld, err = d.openOrCreate()

	return err
}

func (d *DB) openOrCreate() (*gorm.DB, error) {
	db, err := d.open()
	if err == nil {
		return db, nil
	}

	// Error 1049: unable to locate the database
	if !strings.Contains(err.Error(), "Error 1049") {
		return nil, err
	}

	err = d.create()
	if err != nil {
		return nil, err
	}

	return d.open()
}

func (d *DB) open() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      d.cfg.DSN(),
		DefaultStringSize:        256,
		DisableDatetimePrecision: false,
		DontSupportRenameIndex:   true,
		DontSupportRenameColumn:  true,
	}), &gorm.Config{
		Logger: logger.New(logrus.StandardLogger(), logger.Config{
			SlowThreshold:             d.cfg.Logger.SlowThreshold,
			LogLevel:                  d.cfg.Logger.GormLogLevel(),
			Colorful:                  d.cfg.Logger.Colorfule,
			IgnoreRecordNotFoundError: d.cfg.Logger.IgnoreRecordNotFoundError,
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("connect to database [%s] - %w", d.cfg.DSN(), err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get sql database [%s] - %w", d.cfg.DSN(), err)
	}

	sqlDB.SetMaxOpenConns(d.cfg.MaxOpenConnection)
	sqlDB.SetMaxIdleConns(d.cfg.MaxIdleConnection)
	sqlDB.SetConnMaxLifetime(d.cfg.MaxLifeTime)
	sqlDB.SetConnMaxIdleTime(d.cfg.MaxIdleTime)

	var connectionID int
	tx := db.Raw("SELECT CONNECTION_ID()").Scan(&connectionID)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, fmt.Errorf("ping database[%s] - %w", d.cfg.DSN(), err)
	}

	logrus.Debugf("[PING] connected to MySQL database with connection id: %d", connectionID)

	return db, nil
}

func (d *DB) create() error {
	db, err := gorm.Open(mysql.Open(d.cfg.MigrationDSN()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("can't connect to database: [%s] - %w", d.cfg.MigrationDSN(), err)
	}

	defer close(db)

	logrus.Warnf("database doesnot exist, creating it ... %s", d.cfg.MigrationDSN())

	stmt := fmt.Sprintf("CREATE TABLE `%s` CHARACTER SET %s COLLATE %s;", d.cfg.DBName, d.cfg.Charset, d.cfg.Collation)
	if err = db.Exec(stmt).Error; err != nil {
		return fmt.Errorf("can't create database: %w", err)
	}

	return nil
}

func (d *DB) Close() {
	close(d.Realworld)
}

func close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Error(err)
	}

	if err = sqlDB.Close(); err != nil {
		logrus.Error(err)
	}
}
