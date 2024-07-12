package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Host        string `env:"HOST"`
	Port        string `env:"PORT"`
	PGHost      string `env:"POSTGRES_HOST"`
	PGPort      string `env:"POSTGRES_PORT"`
	PGDatabase  string `env:"POSTGRES_DB"`
	PGUser      string `env:"POSTGRES_USER"`
	PGPassword  string `env:"POSTGRES_PASSWORD"`
	RmqHost     string `env:"PABBITMQ_HOST"`
	RmqPort     string `env:"PABBITMQ_PORT"`
	RmqUser     string `env:"PABBITMQ_USER"`
	RmqPassword string `env:"PABBITMQ_PASSWORD"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
		cfg.Host = os.Getenv("HOST")
		cfg.Port = os.Getenv("PORT")
		cfg.PGHost = os.Getenv("POSTGRES_HOST")
		cfg.PGPort = os.Getenv("POSTGRES_PORT")
		cfg.PGDatabase = os.Getenv("POSTGRES_DB")
		cfg.PGUser = os.Getenv("POSTGRES_USER")
		cfg.PGPassword = os.Getenv("POSTGRES_PASSWORD")
		cfg.RmqHost = os.Getenv("PABBITMQ_HOST")
		cfg.RmqPort = os.Getenv("PABBITMQ_PORT")
		cfg.RmqUser = os.Getenv("PABBITMQ_USER")
		cfg.RmqPassword = os.Getenv("PABBITMQ_PASSWORD")
	})
	return &cfg
}
