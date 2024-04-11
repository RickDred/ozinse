// internal/movies/transport/movie_handler.go

package transport

import (
	"net/http"

	"github.com/RickDred/ozinse/internal/movies"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieService movies.MovieServiceInterface
}

func NewMovieHandler(movieService movies.MovieServiceInterface) movies.MovieHandlerInterface {
	return &MovieHandler{
		movieService: movieService,
	}
}

// GetMovies handles the HTTP request to get all movies.
func (h *MovieHandler) GetMovies(c *gin.Context) {
	movies, err := h.movieService.GetMovies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// GetMovie handles the HTTP request to get a single movie by ID.
func (h *MovieHandler) GetMovie(c *gin.Context) {
	movieID := c.Param("id")
	movie, err := h.movieService.GetMovieByID(c.Request.Context(), movieID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// SearchMovies handles the HTTP request to search for movies.
func (h *MovieHandler) SearchMovies(c *gin.Context) {
	query := c.Request.URL.Query()
	movies, err := h.movieService.SearchMovies(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetFavorites(c *gin.Context) {
	// err := h.movieService.AddToFavorites(c.Request.Context(), )
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"message": "Movie added to favorites successfully"})
}

// AddToFavorites handles the HTTP request to add a movie to favorites.
func (h *MovieHandler) AddToFavorites(c *gin.Context) {
	// movieID := c.Param("movieID")
	// err := h.movieService.AddToFavorites(c.Request.Context(), userID, movieID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"message": "Movie added to favorites successfully"})
}
