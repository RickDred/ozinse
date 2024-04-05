// internal/movies/transport/movie_handler.go

package transport

import (
	"net/http"

	"github.com/RickDred/ozinse/internal/movies/service"
	"github.com/gin-gonic/gin"
)

// MovieHandlerImpl implements the MovieHandler interface.
type MovieHandlerImpl struct {
	movieService service.MovieService
}

// NewMovieHandler creates a new instance of MovieHandlerImpl.
func NewMovieHandler(movieService service.MovieService) *MovieHandlerImpl {
	return &MovieHandlerImpl{
		movieService: movieService,
	}
}

// GetMovies handles the HTTP request to get all movies.
func (h *MovieHandlerImpl) GetMovies(c *gin.Context) {
	movies, err := h.movieService.GetMovies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// GetMovie handles the HTTP request to get a single movie by ID.
func (h *MovieHandlerImpl) GetMovie(c *gin.Context) {
	movieID := c.Param("id")
	movie, err := h.movieService.GetMovieByID(c.Request.Context(), movieID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// SearchMovies handles the HTTP request to search for movies.
func (h *MovieHandlerImpl) SearchMovies(c *gin.Context) {
	query := c.Query("q")
	movies, err := h.movieService.SearchMovies(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// AddToFavorites handles the HTTP request to add a movie to favorites.
func (h *MovieHandlerImpl) AddToFavorites(c *gin.Context) {
	userID := c.Param("userID")
	movieID := c.Param("movieID")
	err := h.movieService.AddToFavorites(c.Request.Context(), userID, movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie added to favorites successfully"})
}
