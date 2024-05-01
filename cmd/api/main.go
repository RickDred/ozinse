package main

import (
	"github.com/RickDred/ozinse/config"
	_ "github.com/RickDred/ozinse/docs"
	"github.com/RickDred/ozinse/internal/api"
	"github.com/RickDred/ozinse/pkg/db"
	"github.com/gin-gonic/gin"
)

// @title Ozinse API
// @version         1.0
// @description     Api for ozinse application.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:3000
// @BasePath  /
// @schemes   http

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @securityDefinitions.ApiKey ApiKeyAuth
// @in header
// @name Authorization
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
	router := gin.New()

	app.Start(router)
}
