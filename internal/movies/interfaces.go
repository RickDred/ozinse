package movies

import (
	"context"
	"net/url"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
)

// MovieHandlerInterface defines the methods implemented by the movie handlers.
type MovieHandlerInterface interface {
	GetMovies(*gin.Context)
	GetMovie(*gin.Context)
	CreateMovie(*gin.Context)
	EditMovie(*gin.Context)
	DeleteMovie(*gin.Context)
	GetMovieSeries(*gin.Context)
	GetMoviesByCategory(*gin.Context)

	SearchMovies(*gin.Context)
	GetFavorites(*gin.Context)
	AddToFavorites(*gin.Context)
}

// MovieServiceInterface defines the methods implemented by the movie service.
type MovieServiceInterface interface {
	GetMovieByID(context.Context, uint) (*models.Movie, error)
	GetMovies(context.Context) ([]models.Movie, error)
	CreateMovie(context.Context, *models.Movie) (*models.Movie, error)
	EditMovie(context.Context, uint, *models.Movie) (*models.Movie, error)
	DeleteMovie(context.Context, uint) error
	// GetMovieSeries(context.Context)
	GetMoviesByCategory(context.Context, string) ([]models.Movie, error)

	SearchMovies(context.Context, url.Values) ([]models.Movie, error)
	AddToFavorites(context.Context, string, string) error
}

// MovieRepositoryInterface defines the methods implemented by the movie repository.
type MovieRepositoryInterface interface {
	GetByID(context.Context, uint) (*models.Movie, error)
	GetAll(context.Context) ([]models.Movie, error)
	Insert(context.Context, *models.Movie) (*models.Movie, error)
	Update(context.Context, uint, *models.Movie) (*models.Movie, error)
	Delete(context.Context, uint) error
	GetAllByCategory(context.Context, string) ([]models.Movie, error)

	Search(context.Context, map[string]interface{}) ([]models.Movie, error)
	AddToFavorites(context.Context, *models.User, *models.Movie) error
}
