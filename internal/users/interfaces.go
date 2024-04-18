package users

import (
	"github.com/RickDred/ozinse/internal/models"
	"github.com/gin-gonic/gin"
)

type UsersHandlerInterface interface {
	GetProfile(*gin.Context)
	EditProfile(*gin.Context)
	ChangePassword(*gin.Context)
}

type UsersServiceInterface interface {
	GetProfile(uint) (*models.User, error)
	EditProfile(user *models.User, name string, phone int) error
	ChangePassword(*models.User, string, string) error
}

type UsersRepositoryInterface interface {
	GetUser(uint) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(uint) error
}
