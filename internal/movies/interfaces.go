package movies

import (
	"context"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
)

// MovieHandlerInterface defines the methods implemented by the movie handlers.
type MovieHandlerInterface interface {
	GetMovies(*gin.Context)
	GetMovie(*gin.Context)
	SearchMovies(*gin.Context)
	AddToFavorites(*gin.Context)
}

// MovieServiceInterface defines the methods implemented by the movie service.
type MovieServiceInterface interface {
	GetMovieByID(context.Context, string) (*models.Movie, error)
	GetMovies(context.Context) ([]*models.Movie, error)
	SearchMovies(context.Context, string) ([]*models.Movie, error)
	AddToFavorites(context.Context, string, string) error
}

// MovieRepositoryInterface defines the methods implemented by the movie repository.
type MovieRepositoryInterface interface {
	GetByID(context.Context, string) (*models.Movie, error)
	GetAll(context.Context) ([]*models.Movie, error)
	Search(context.Context, string) ([]*models.Movie, error)
	AddToFavorites(context.Context, string, string) error
}
