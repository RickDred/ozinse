package service

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/RickDred/ozinse/config"
	"github.com/RickDred/ozinse/internal/auth"
	"github.com/RickDred/ozinse/internal/models"
	"github.com/RickDred/ozinse/pkg/smtp"
)

type service struct {
	repo     auth.RepoInterface
	jwtCfg   config.JWTConfig
	emailCfg config.EmailConfig
}

func New(repo auth.RepoInterface, jwtCfg config.JWTConfig, emailCfg config.EmailConfig) auth.ServiceInterface {
	return &service{
		repo:     repo,
		jwtCfg:   jwtCfg,
		emailCfg: emailCfg,
	}
}

func (s *service) Register(ctx context.Context, user *models.User, repeatedPassword string) (string, error) {
	user.Standardize()

	repeatedPassword = strings.TrimSpace(repeatedPassword)

	if repeatedPassword != user.Password {
		return "", models.ErrPasswrodsNotMatch
	}

	if err := user.Validate(user.ValidateEmail, user.ValidatePassword); err != nil {
		return "", err
	}

	if err := user.HashPassword(); err != nil {
		return "", err
	}

	if uu, err := s.repo.GetByEmail(ctx, user.Email); uu != nil || err == nil {
		return "", models.ErrEmailExists
	}

	_, err := s.repo.Insert(ctx, user)
	if err != nil {
		return "", err
	}

	user.CleanPassword()

	token, err := s.generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) Login(ctx context.Context, input *models.User) (string, error) {
	user, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	if err := user.ComparePassword(input.Password); err != nil {
		return "", err
	}

	user.CleanPassword()

	token, err := s.generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) PasswordRecover(ctx context.Context, user *models.User) (bool, error) {
	exUser, err := s.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		return false, err
	}

	// generate new password for user
	newPassword := GeneratePassword(10)
	user.Password = newPassword

	if err := user.HashPassword(); err != nil {
		return false, err
	}

	if err := s.repo.PasswordRecover(ctx, exUser.ID, user.Password); err != nil {
		return false, err
	}

	// send email with new password

	err = smtp.SendEmail(
		s.emailCfg.Identity,
		s.emailCfg.Username,
		s.emailCfg.Password,
		s.emailCfg.Host,
		s.emailCfg.Addr,
		"ozinse administartion",
		user.Email,
		"Password recovery",
		"Your new password: "+newPassword,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GeneratePassword(length int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+"

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}
