package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var C *Config

type Config struct {
	Application string     `yaml:"application"`
	HTTPServer  HTTPServer `mapstructure:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address"`
}

func Init(cfgFile string) error {
	v := viper.New()
	v.SetConfigFile(cfgFile)
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.SetEnvPrefix("linky")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Error in reading config file, %v", err)

		return err
	}

	err = v.Unmarshal(&C)
	if err != nil {
		fmt.Println("Error in unmarshalling config file")

		return err
	}

	return nil
}
