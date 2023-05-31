package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// ImportConfigs find config file and creates a new Config struct
func ImportConfigs() (config *Config) {
	viper.SetConfigFile("../config.json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(err)
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse configuration file into structure 'Config'. Error: %s", err.Error()))
	}

	return
}
