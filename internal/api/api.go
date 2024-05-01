package api

import (
	"fmt"

	"github.com/RickDred/ozinse/config"
	_ "github.com/RickDred/ozinse/docs"
	"github.com/RickDred/ozinse/internal/auth"
	arepo "github.com/RickDred/ozinse/internal/auth/repository"
	aservice "github.com/RickDred/ozinse/internal/auth/service"
	atransport "github.com/RickDred/ozinse/internal/auth/transport"
	"github.com/RickDred/ozinse/internal/models"
	"github.com/RickDred/ozinse/internal/movies"
	mrepo "github.com/RickDred/ozinse/internal/movies/repository"
	mservice "github.com/RickDred/ozinse/internal/movies/service"
	mtransport "github.com/RickDred/ozinse/internal/movies/transport"
	"github.com/RickDred/ozinse/internal/users"
	urepo "github.com/RickDred/ozinse/internal/users/repository"
	uservice "github.com/RickDred/ozinse/internal/users/service"
	utransport "github.com/RickDred/ozinse/internal/users/transport"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Api struct {
	DB  *gorm.DB
	Cfg *config.Config
}

func (a *Api) Start(router *gin.Engine) {
	if err := a.DB.AutoMigrate(&models.User{}, &models.Movie{}, &models.Video{}, &models.Category{}, &models.Genre{}); err != nil {
		panic(err)
	}
	//router := gin.New()
	router.Use(gin.Logger())

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// need to move into other function
	authrepo := arepo.New(a.DB)
	authservice := aservice.New(authrepo, a.Cfg.JWT, a.Cfg.Email)
	authhandlers := atransport.New(authservice)
	authGorup := router.Group("/auth")
	auth.SetRoutes(authGorup, authhandlers)

	router.Use(authhandlers.Authorize())

	moviesrepo := mrepo.NewMovieRepository(a.DB)
	mserv := mservice.NewMovieService(moviesrepo)
	mhandl := mtransport.NewMovieHandler(mserv)
	moviesGroup := router.Group("/movies")
	movies.InitRoutes(moviesGroup, mhandl)

	usersrepo := urepo.NewUsersRepository(a.DB)
	usersserv := uservice.NewService(usersrepo)
	usershandl := utransport.NewHandler(usersserv)
	usersGroup := router.Group("/users")
	users.InitRoutes(usersGroup, usershandl)

	addr := fmt.Sprintf("%v:%v", a.Cfg.Host, a.Cfg.Port)

	// authservice.Register(context.Background(), &models.User{
	// 	Email:    "admin@admin.com",
	// 	Password: "123456",
	// 	Role:     "admin",
	// 	Name:     "admin",
	// 	Phone:    87784614283,
	// })

	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
