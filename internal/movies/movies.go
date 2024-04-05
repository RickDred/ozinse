// internal/movies/movies.go

package movies

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, handlers MovieHandler) {
	router.GET("/", handlers.GetMovies)
	router.GET("/:id", handlers.GetMovie)
	router.GET("/search", handlers.SearchMovies)
	router.POST("/:userID/favorites/:movieID", handlers.AddToFavorites)
}
