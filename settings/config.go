package settings

import (
	"my/v1/model"

	"github.com/spf13/viper"
)

// GetConfig to parse configuration
func GetConfig() (*model.Config, error) {
	return GetConfigFromFile("default")
}

// GetConfigFromFile to parse configuration from file
func GetConfigFromFile(fileName string) (*model.Config, error) {
	if fileName == "" {
		fileName = "default"
	}
	viper.SetConfigName(fileName)
	viper.AddConfigPath("../conf/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		// Find and read the config file
		return nil, err
	}
	config := &model.Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
