// internal/movies/movies.go

package movies

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, handlers MovieHandlerInterface) {
	router.GET("/", handlers.GetMovies)
	router.GET("/:id", handlers.GetMovie)
	router.POST("/", handlers.CreateMovie)
	router.PUT("/:id", handlers.EditMovie)
	router.DELETE("/:id", handlers.DeleteMovie)
	router.GET("/series/:id", handlers.GetMovieSeries)
	router.GET("/categories/:name", handlers.GetMoviesByCategory)
	router.POST("/series", handlers.UploadVideo)

	// wait a minute
	router.GET("/search", handlers.SearchMovies)
	router.GET("/favorites", handlers.GetFavorites)
	router.POST("/favorites/:movieID", handlers.AddToFavorites)
}
