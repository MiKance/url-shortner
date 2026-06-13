package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mikance/url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoadEnv()
	m, err := migrate.New("file://migrations/", cfg.Postgres.ConnString())
	if err != nil {
		log.Fatal("cannot connect to database or open migrations files:", err)
	}
	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no change in database schemas")
			return
		}
		log.Fatal("cannot use migrations files up: ", err)
	}
	fmt.Println("database migrations successfully applied")
}
