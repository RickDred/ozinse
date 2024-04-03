package app

import (
	"github.com/RickDred/ozinse/internal/auth"
	"github.com/RickDred/ozinse/internal/auth/repository"
	"github.com/RickDred/ozinse/internal/auth/service"
	"github.com/RickDred/ozinse/internal/auth/transport"
	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	DB   *gorm.DB
	Port string
}

func (a *App) Start() {
	if err := a.DB.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}
	router := gin.New()
	router.Use(gin.Logger())

	// need to move into other function
	repo := repository.New(a.DB)
	service := service.New(repo)
	handlers := transport.New(service)
	authGorup := router.Group("/auth")
	auth.SetRoutes(authGorup, handlers)

	router.Run(a.Port)
}
