package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/wackGarcia/books/data/db"
)

type Config struct {
	HttpServer  int
	Environment string
	GoogleAPI   string
	Storage     db.Storage
}

func (config *Config) LoadConfig() error {
	v := viper.New()
	envPrefix := "books"

	v.SetConfigName("books")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	for _, key := range v.AllKeys() {
		keyReplaced := strings.ReplaceAll(strings.ReplaceAll(key, "_", "."), envPrefix+".", "")
		v.BindEnv(key, keyReplaced)
		v.SetDefault(keyReplaced, v.Get(key))
	}

	if err := v.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}
