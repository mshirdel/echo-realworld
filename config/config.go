package config

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

type Config struct {
	Application string     `mapstructure:"application"`
	HTTPServer  HTTPServer `mapstructure:"http_server"`
	Logging     Logging    `mapstructure:"logging"`
	Database    Database   `mapstructure:"database"`
}

type HTTPServer struct {
	Address      string        `mapstructure:"address"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

type Logging struct {
	Level string `mapstructure:"level"`
}

type Database struct {
	Host              string        `mapstructure:"host"`
	Port              int           `mapstructure:"port"`
	User              string        `mapstructure:"user"`
	Password          string        `mapstructure:"password"`
	DBName            string        `mapstructure:"dbname"`
	Charset           string        `mapstructure:"charset"`
	Collation         string        `mapstructure:"collation"`
	ParseTime         bool          `mapstructure:"parse_time"`
	Location          string        `mapstructure:"location"`
	MaxLifeTime       time.Duration `mapstructure:"max_life_time"`
	MaxIdleTime       time.Duration `mapstructure:"max_idel_time"`
	MaxOpenConnection int           `mapstructure:"max_open_connection"`
	MaxIdleConnection int           `mapstructure:"max_idel_connection"`
	Logger            Logger        `mapstructure:"logger"`
}

func (d *Database) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t&charset=%s&collation=%s&loc=%s&multiStatements=true",
		d.User, d.Password, d.Host, d.Port, d.DBName, d.ParseTime, d.Charset, d.Collation, url.PathEscape(d.Location))
}

func (d *Database) MigrationDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t&charset=%s&collation=%s&loc=%s&multiStatements=true",
		d.User, d.Password, d.Host, d.Port, "", d.ParseTime, d.Charset, d.Collation, url.PathEscape(d.Location))
}

type Logger struct {
	SlowThreshold             time.Duration `mapstructure:"slow_threshold"`
	Level                     string        `mapstructure:"level"`
	Colorfule                 bool          `mapstructure:"colorfule"`
	IgnoreRecordNotFoundError bool          `mapstructure:"ignore_record_not_found_error"`
}

func (l Logger) GormLogLevel() logger.LogLevel {
	switch l.Level {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn", "warning":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Warn
	}
}

var _defulatConfig = ""
var prifix = "realworld"

func Init(cfgFile string) (*Config, error) {
	var c Config

	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewReader([]byte(_defulatConfig))); err != nil {
		return nil, fmt.Errorf("error loading default config: %w", err)
	}

	v.SetConfigFile(cfgFile)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.SetEnvPrefix(prifix)
	v.AutomaticEnv()

	switch err := v.MergeInConfig(); err.(type) {
	case nil:
	case *os.PathError:
		logrus.Infof("config file %s not found", cfgFile)
	default:
		logrus.Warnf("faild to load config file: %s", err)
	}

	if err := v.UnmarshalExact(&c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &c, nil
}
