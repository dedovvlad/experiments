package main

import "github.com/vrischmann/envconfig"

type (
	Config struct {
		DB       string `envconfig:"default=host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable"`
		FilePath string `envconfig:"default=work_with_database/passports_demo.csv"`
	}
)

func InitConfig() *Config {
	config := &Config{}
	if err := envconfig.Init(config); err != nil {
		return nil
	}
	return config
}
