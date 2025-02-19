package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	// env: считываем из переменной окружения, если файла конфига нет, то приложение не запустится
	Env        string `yaml:"env" env:"ENV" env-default:"local" env-required:"true"`
	Storage    `yaml:"storage"`
	HTTPServer `yaml:"http_server"`
}

type Storage struct {
	DBHost          string `yaml:"db_host" env_required:"true"`
	DBPort          string `yaml:"db_port" env_required:"true"`
	DBUser          string `yaml:"db_user" env_required:"true"`
	DBName          string `yaml:"db_name" env_required:"true"`
	DBPass          string `yaml:"db_pass" env_required:"true" env:"DB_SERVER_PASSWORD"`
	DSN             string `yaml:"db_dsn" env_required:"true"`
	DBMaxConnection int32  `yaml:"db_max_connection"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env_required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func NewConfig() *Config {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// проверяем наличие файла конфигурации
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file do not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("can`t read config: %s", err)
	}

	return &cfg

}
