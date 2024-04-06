package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/RickDred/ozinse/internal/movies"
	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) movies.MovieRepositoryInterface {
	return &MovieRepository{db}
}

func (r *MovieRepository) GetByID(ctx context.Context, id string) (*models.Movie, error) {
	var movie models.Movie
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&movie)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("movie with id %s not found", id)
		}
		return nil, result.Error
	}
	return &movie, nil
}

func (r *MovieRepository) GetAll(ctx context.Context) ([]models.Movie, error) {
	var movies []models.Movie
	result := r.db.WithContext(ctx).Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}

// Search searches for movies in the database based on a query string.
func (r *MovieRepository) Search(ctx context.Context, criteria map[string]interface{}) ([]models.Movie, error) {
	var movies []models.Movie

	query := r.db.WithContext(ctx)
	for key, value := range criteria {
		switch key {
		case "title":
			query = query.Where("title LIKE ?", "%"+value.(string)+"%")
		case "producer":
			query = query.Where("producer LIKE ?", "%"+value.(string)+"%")
		case "director":
			query = query.Where("director LIKE ?", "%"+value.(string)+"%")
		case "year":
			query = query.Where("year = ?", value.(string))
		case "type":
			query = query.Where("type = ?", value.(string))
		case "tag":
			query = query.Where("tags @> ARRAY[?]::text[]", value.(string))
		case "tags":
			tags := value.([]string)
			for _, tag := range tags {
				query = query.Where("tags @> ARRAY[?]::text[]", tag)
			}
		case "seasons":
			query = query.Where("seasons = ?", value.(int))
		}
	}

	if err := query.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MovieRepository) AddToFavorites(ctx context.Context, user *models.User, movie *models.Movie) error {
	var existingMovie models.Movie
	if err := r.db.Model(user).Where("id = ?", movie.ID).Association("Favorites").Find(&existingMovie); err == nil {
		// Movie already exists in favorites
		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// An unexpected error occurred
		return err
	}

	// Add movie to favorites
	if err := r.db.Model(user).Association("Favorites").Append(movie); err != nil {
		return err
	}

	return nil
}
