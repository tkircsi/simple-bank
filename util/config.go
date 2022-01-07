package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string   `mapstructure:"DB_DRIVER"`
	DBSource      string   `mapstructure:"DB_SOURCE"`
	ServerAddress string   `mapstructure:"SERVER_ADDRESS"`
	Currency      []string `mapstructure:"CURRENCY"`
}

var AppConfig Config

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return nil, err
	}
	return &AppConfig, nil
}
