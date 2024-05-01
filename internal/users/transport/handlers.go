package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get user by ID
// @Description get string by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (h *UsersHandler) GetProfile(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := h.service.GetProfile(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UsersHandler) EditProfile(c *gin.Context) {
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	var input struct {
		Name      string `json:"name"`
		Phone     int    `json:"phone"`
		Birthdate string `json:"birthdate"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	err = h.service.EditProfile(user, input.Name, input.Phone, input.Birthdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "profile updated successfully"})
}

func (h *UsersHandler) ChangePassword(c *gin.Context) {
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	err = h.service.ChangePassword(user, input.OldPassword, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})
}
