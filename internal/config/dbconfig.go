package config

import (
	"fmt"

	"github.com/RickDred/ozinse/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func (cfg *DBConfig) InitDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db
}
