package main

import (
	"github.com/RickDred/ozinse/config"
	"github.com/RickDred/ozinse/internal/api"
	"github.com/RickDred/ozinse/pkg/db"
)


// @title Ozinse API
func main() {
	cfg := config.NewConfig()

	db, err := db.Connect(cfg.Postgres.DSN())
	if err != nil {
		panic(err)
	}

	app := api.Api{
		DB:  db,
		Cfg: cfg,
	}

	app.Start()
}
