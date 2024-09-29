package config

import (
	"errors"
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

	if err := checkVars(); err != nil {
		logger.Error(err.Error(), constants.ConfigCategory)

		return err
	}

	return nil
}

func checkVars() error {
	if ServerConfig.ApiPort == 0 || ServerConfig.ApiEntry == "" {
		return errors.New(constants.EmptyConfigVarError.Error() + "API_PORT, API_ENTRY must be set")
	}

	if ServerConfig.PostgresDSN == "" {
		return errors.New(constants.EmptyConfigVarError.Error() + "POSTGRESQL_DSN must be set")
	}

	if ServerConfig.SpotifyId == "" || ServerConfig.SpotifySecret == "" {
		return errors.New(constants.EmptyConfigVarError.Error() + "SPOTIFY_CLIENT_ID, SPOTIFY_SECRET must be set")
	}

	return nil
}
