package config

import (
	"final-project/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbInit() *gorm.DB {
	env := godotenv.Load(".env")

	if env != nil {
		log.Fatal("Failed to load env")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbPort := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, dbPort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database")
	}

	db.AutoMigrate(models.User{}, models.Photo{})

	if err != nil {
		errorMessage := fmt.Sprintf("Failed to migrate models: %s", err.Error())
		log.Fatal(errorMessage)
	}

	return db
}
