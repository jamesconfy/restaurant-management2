package config

import "github.com/spf13/viper"

type Env struct {
	MODE             string `mapstructure:"MODE"`
	ADDR             string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN string `mapstructure:"SECRET_KEY_TOKEN"`
	CACHE            bool   `mapstructure:"CACHE"`
	HOST             string `mapstructure:"HOST"`
	PORT             string `mapstructure:"PORT"`
	PASSWD           string `mapstructure:"PASSWD"`
	EMAIL            string `mapstructure:"EMAIL"`
	EXPIRES_AT       int64  `mapstructure:"EXPIRES_AT"`

	POSTGRES_HOST      string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_USER      string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD  string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DB        string `mapstructure:"POSTGRES_DB"`
	DATABASE_SOURCE_DB string `mapstructure:"DATABASE_SOURCE_DB"`

	REDIS_USERNAME string `mapstructure:"REDIS_USERNAME"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
	REDIS_ADDRESS  string `mapstructure:"REDIS_ADDRESS"`
}

var Environment *Env

func init() {
	viper.AutomaticEnv()

	Environment = &Env{
		MODE:               viper.GetString("MODE"),
		ADDR:               viper.GetString("ADDR"),
		SECRET_KEY_TOKEN:   viper.GetString("SECRET_KEY_TOKEN"),
		CACHE:              viper.GetBool("CACHE"),
		HOST:               viper.GetString("HOST"),
		PORT:               viper.GetString("PORT"),
		PASSWD:             viper.GetString("PASSWD"),
		EMAIL:              viper.GetString("EMAIL"),
		EXPIRES_AT:         viper.GetInt64("EXPIRES_AT"),
		POSTGRES_HOST:      viper.GetString("POSTGRES_HOST"),
		POSTGRES_USER:      viper.GetString("POSTGRES_USER"),
		POSTGRES_PASSWORD:  viper.GetString("POSTGRES_PASSWORD"),
		POSTGRES_DB:        viper.GetString("POSTGRES_DB"),
		DATABASE_SOURCE_DB: viper.GetString("DATABASE_SOURCE_DB"),
		REDIS_USERNAME:     viper.GetString("REDIS_USERNAME"),
		REDIS_PASSWORD:     viper.GetString("REDIS_PASSWORD"),
		REDIS_ADDRESS:      viper.GetString("REDIS_ADDRESS"),
	}
}
