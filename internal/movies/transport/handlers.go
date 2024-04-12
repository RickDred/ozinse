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
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie, err := h.movieService.GetMovieByID(c.Request.Context(), uint(movieID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
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

///// Wait a minute

// SearchMovies handles the HTTP request to search for movies.
func (h *MovieHandler) SearchMovies(c *gin.Context) {
	// query := c.Request.URL.Query()
	// movies, err := h.movieService.SearchMovies(c.Request.Context(), query)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, movies)
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
