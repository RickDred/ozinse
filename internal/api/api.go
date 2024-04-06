package api

import (
	"fmt"

	"github.com/RickDred/ozinse/internal/auth"
	"github.com/RickDred/ozinse/internal/auth/repository"
	"github.com/RickDred/ozinse/internal/auth/service"
	"github.com/RickDred/ozinse/internal/auth/transport"
	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB   *gorm.DB
	Port int
	Host string
}

func (a *Server) Start() {
	if err := a.DB.AutoMigrate(&models.User{}, &models.Movie{}, &models.Video{}); err != nil {
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

	addr := fmt.Sprintf("%v:%v", a.Host, a.Port)

	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
