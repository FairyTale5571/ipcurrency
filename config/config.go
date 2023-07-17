package config

import (
	"github.com/spf13/viper"
)

func ReadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
