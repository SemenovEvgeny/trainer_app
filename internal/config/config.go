package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `env:"ENV"          env-default:"local" env-required:"true" json:"env"`
	Storage    `json:"storage"`
	HTTPServer `json:"http_server"`
}

type Storage struct {
	Host          string `env_required:"true"          json:"host"`
	Port          string `env_required:"true"          json:"port"`
	User          string `env_required:"true"          json:"user"`
	Name          string `env_required:"true"          json:"name"`
	Pass          string `env_required:"true"          json:"pass"          env:"DB_SERVER_PASSWORD"`
	DSN           string `env_required:"true"          json:"dsn"`
	MaxConnection int32  `json:"max_connection"`
}

type HTTPServer struct {
	Address     string        `json:"address"`
	Port        int           `json:"port"`
	Timeout     time.Duration `env-default:"4s"    json:"timeout"`
	IdleTimeout time.Duration `env-default:"60s"   json:"idle_timeout"`
	User        string        `env-required:"true" json:"user"`
	Password    string        `env_required:"true" json:"password"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			log.Fatal("CONFIG_PATH is not set")
		}

		// проверяем наличие файла конфигурации
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			log.Fatalf("config file do not exist: %s", configPath)
		}

		var cfg Config
		if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
			log.Fatalf("can't read config: %s", err)
		}
		instance = &cfg
	})

	return instance

}
