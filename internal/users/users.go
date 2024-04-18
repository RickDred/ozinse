package users

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.RouterGroup, handlers UsersHandlerInterface) {
	router.GET("/:id", handlers.GetProfile)
	router.PUT("/profile", handlers.EditProfile)
	router.PUT("/profile/changePassword", handlers.ChangePassword)
}
