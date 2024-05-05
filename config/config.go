package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Application string     `mapstructure:"application"`
	HTTPServer  HTTPServer `mapstructure:"http_server"`
	Logging     Logging    `mapstructure:"logging"`
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
