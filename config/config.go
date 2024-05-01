// config/config.go

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     int
	Host     string
	Postgres PostgresConfig
	JWT      JWTConfig
	Email    EmailConfig
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	SecretKey string
	Expire    int
}

type EmailConfig struct {
	Host     string
	Username string
	Password string
	Identity string
	Addr     string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		port = 8080
	}
	postgresPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		postgresPort = 5432
	}

	expire, err := strconv.Atoi(getEnv("JWT_EXPIRE", "60"))
	if err != nil {
		expire = 60
	}

	cfg := &Config{
		Port: port,
		Host: getEnv("SERVER_HOST", "localhost"),
		Postgres: PostgresConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     postgresPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "ozinse"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey: getEnv("JWT_SECRET_KEY", "secret"),
			Expire:    expire,
		},
		Email: EmailConfig{
			Host:     getEnv("EMAIL_HOST", "smtp.example.com"),
			Username: getEnv("EMAIL_USERNAME", "username@gmail.com"),
			Password: getEnv("EMAIL_PASSWORD", "password"),
			Identity: getEnv("EMAIL_IDENTITY", ""),
			Addr:    getEnv("EMAIL_ADDR", "smtp.example.com:587"),
		},
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return value
}

func (pc *PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pc.Host, pc.Port, pc.User, pc.Password, pc.DBName, pc.SSLMode)
}
