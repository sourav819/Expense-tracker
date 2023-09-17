package config

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig() (Config, error) {

	var (
		l   = logrus.New()
		cfg Config
	)
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	//to start reading config values
	err := viper.ReadInConfig()
	if err != nil {
		// return err
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			l.Fatal("Error reading config file: ", err)
		}
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	if cfg.Server.Port == "" {
		return cfg, errors.New("error reading config")
	}
	return cfg, nil

}
