package transport

import (
	"net/http"
	"strconv"

	"github.com/RickDred/ozinse/internal/helpers"
	"github.com/RickDred/ozinse/internal/models"
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
	user, err := helpers.GetUserFromGin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie, isFav, sameMovies, err := h.movieService.GetMovieByID(c.Request.Context(), user, uint(movieID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movie":       movie,
		"is_favorite": isFav,
		"same_movies": sameMovies,
	})
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	user, err := helpers.GetUserFromGin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	if user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you have not permission"})
		return
	}

	var movie models.Movie
	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMovie, err := h.movieService.CreateMovie(c.Request.Context(), &movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMovie)
}

func (h *MovieHandler) EditMovie(c *gin.Context) {
	user, err := helpers.GetUserFromGin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	if user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you have not permission"})
		return
	}

	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var updatedMovie models.Movie
	if err := c.BindJSON(&updatedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	editedMovie, err := h.movieService.EditMovie(c.Request.Context(), uint(movieID), &updatedMovie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, editedMovie)
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	user, err := helpers.GetUserFromGin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	if user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you have not permission"})
		return
	}

	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.movieService.DeleteMovie(c.Request.Context(), uint(movieID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}

func (h *MovieHandler) GetMovieSeries(c *gin.Context) {
	movieID := c.Param("id")

	// Parse movie ID to uint
	id, err := strconv.ParseUint(movieID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	// Call the service layer to get movie series
	videos, err := h.movieService.GetMovieSeries(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}

func (h *MovieHandler) GetMoviesByCategory(c *gin.Context) {
	category := c.Param("name")

	movies, err := h.movieService.GetMoviesByCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) SearchMovies(c *gin.Context) {
	var filters models.MoviesFilter

	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movies, err := h.movieService.SearchMovies(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

///// Wait a minute

func (h *MovieHandler) UploadVideo(c *gin.Context) {
}
