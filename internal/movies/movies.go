// internal/movies/movies.go

package movies

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, handlers MovieHandlerInterface) {
	router.GET("/", handlers.GetMovies)
	router.GET("/:id", handlers.GetMovie)
	router.GET("/search", handlers.SearchMovies)
	router.GET("/favorites", handlers.GetFavorites)
	router.POST("/favorites/:movieID", handlers.AddToFavorites)
}
