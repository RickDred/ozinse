package service

import (
	"time"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

func (s *service) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Minute * time.Duration(s.jwtCfg.Expire)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtCfg.SecretKey))
}
