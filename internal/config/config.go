package config

import (
	"log/slog"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8899"`
	} `yaml:"listen"`
}

var instance *Config

var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		slog.Info("Read App congiguration")
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yaml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			slog.Info(help)
			slog.Error(err.Error())
		}
	})

	return instance
}
