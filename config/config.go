package config

import (
	"context"

	"github.com/k0kubun/pp"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Database string
	SSLMode  string
}

type Config struct {
	Name           string
	Port           string
	DatabaseConfig DatabaseConfig
}

func New() *Config {
	var config Config

	viper.AutomaticEnv()
	viper.SetEnvPrefix("WISH")
	viper.SetDefault("PORT", "3449")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USERNAME", "admin")
	viper.SetDefault("DB_PASSWORD", "admin12356")
	viper.SetDefault("DB_DATABASE", "wish")
	viper.SetDefault("DB_SSLMode", "disable")

	config.Name = viper.GetString("WISH")
	config.Port = viper.GetString("PORT")
	config.DatabaseConfig = DatabaseConfig{
		Host:     viper.GetString("DB_HOST"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		Database: viper.GetString("DB_DATABASE"),
		SSLMode:  viper.GetString("DB_SSLMode"),
	}

	return &config
}

func GetDatabaseConfig(config *Config) *DatabaseConfig {
	return &config.DatabaseConfig
}

func Print(lifecycle fx.Lifecycle, c *Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			_, err := pp.Println(c)
			return err
		},
	})

}
