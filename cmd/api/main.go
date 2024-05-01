package main

import (
	_ "github.com/RickDred/ozinse/docs"
	"github.com/RickDred/ozinse/config"
	"github.com/RickDred/ozinse/internal/api"
	"github.com/RickDred/ozinse/pkg/db"
)

// @title Ozinse API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:3000
// @BasePath  /
// @schemes   http

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
