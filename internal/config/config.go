package config

import (
	"fmt"
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Host         string        `env:"SERVER_HOST" envDefault:"localhost"`
	Port         int           `env:"SERVER_PORT" envDefault:"8080"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"2s"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT" envDefault:"20s"`
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	Database string `env:"POSTGRES_DB" envDefault:"url-shortner"`
}

func (p *PostgresConfig) ConnString() string {
	s := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.Database)
	return s
}

type RedisConfig struct {
	Host string        `env:"REDIS_HOST" envDefault:"localhost"`
	Port int           `env:"REDIS_PORT" envDefault:"6379"`
	TTL  time.Duration `env:"REDIS_TTL" envDefault:"1h"`
}

func MustLoadEnv() *Config {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./deployment/.env", cfg)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return cfg
}
