package main

import (
	"log"
	"os"

	"github.com/RickDred/ozinse/config"
	"github.com/RickDred/ozinse/internal/app"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	dbcfg := config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db := dbcfg.InitDB()

	app := app.App{
		DB:   db,
		Port: ":3000",
	}

	app.Start()
}
