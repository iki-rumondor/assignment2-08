package database

import (
	"fmt"
	"log"
	"os"

	"github.com/iki-rumondor/assignment2-GLNG-KS-08-08/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	db *gorm.DB
}

func InitPostgresDb() (*postgresDB, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file")
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := newPostgresDb(connStr)

	if err != nil {
		return nil, err
	}

	return &postgresDB{db: db}, nil

}

func newPostgresDb(connStr string) (*gorm.DB, error) {
	gormDb, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	gormDb.Debug().AutoMigrate(models.Order{}, models.Item{})

	return gormDb, nil
}
