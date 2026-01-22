package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port        int           `mapstructure:"PORT"`
		Env         string        `mapstructure:"ENV"`
		ReadTimeout time.Duration `mapstructure:"READ_TIMEOUT"`
	} `mapstructure:"SERVER"`

	Database struct {
		DSN         string `mapstructure:"DSN"`
		MaxOpenConn int    `mapstructure:"MAX_OPEN_CONN"`
	} `mapstructure:"DATABASE"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
