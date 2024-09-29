package config

import (
	"github.com/spf13/viper"
	"music-library/init/logger"
	"music-library/pkg/constants"
)

var ServerConfig Config

type Config struct {
	ApiPort  int    `mapstructure:"API_PORT"`
	ApiDebug bool   `mapstructure:"API_DEBUG"`
	ApiEntry string `mapstructure:"API_ENTRY"`

	PostgresDSN string `mapstructure:"POSTGRESQL_DSN"`

	SpotifyId     string `mapstructure:"SPOTIFY_ID"`
	SpotifySecret string `mapstructure:"SPOTIFY_SECRET"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Error(err.Error(), constants.ConfigCategory)

		return err
	}

	if err := viper.Unmarshal(&ServerConfig); err != nil {
		logger.Error(err.Error(), constants.ConfigCategory)

		return err
	}

	return nil
}
