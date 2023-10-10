package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"development"`
	HTTPServer `yaml:"http_server"`
	DBConfig `yaml:"db_config"`
}

type HTTPServer struct {
	Host     		string        `yaml:"host" env-default:"127.0.0.1"`
	Port     		string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DBConfig struct {
  Host				string				`yaml:"host" env-default:"127.0.0.1"`
	Database		string				`yaml:"database" env-default:"apiwalk"`
	Username		string				`yaml:"username" env-default:"root"`
	Password		string				`yaml:"password" env-default:""`
}

func MustLoad() *Config {
	// Получаем путь до конфиг-файла из env-переменной CONFIG_PATH
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	// Проверяем существование конфиг-файла
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	var cfg Config

	// Читаем конфиг-файл и заполняем нашу структуру
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
			log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}