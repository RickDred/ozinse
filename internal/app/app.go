package app

import (
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

	repo := repository.New(a.DB)

	s := service.New(repo)

	h := transport.New(s)

	r := gin.Default()

	r.POST("/register", h.Register)

	r.Run(a.Port)
}
