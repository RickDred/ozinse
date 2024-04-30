package auth

import (
	"context"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
)

type HandlersInterface interface {
	Register(*gin.Context)
	Login(*gin.Context)
	PasswordRecover(*gin.Context)

	Authorize() gin.HandlerFunc
}

type ServiceInterface interface {
	Register(context.Context, *models.User, string) (string, error)
	Login(context.Context, *models.User) (string, error)
	PasswordRecover(context.Context, *models.User) (bool, error)
}

type RepoInterface interface {
	Insert(context.Context, *models.User) (uint, error)
	GetByEmail(context.Context, string) (*models.User, error)
	GetByID(context.Context, uint) (*models.User, error)
	GetAll(context.Context) ([]models.User, error)
	PasswordRecover(context.Context, *models.User) error
}
