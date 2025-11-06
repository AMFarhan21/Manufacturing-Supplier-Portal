package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error on loading .env")
	}

	dsn := os.Getenv("CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Error on connecting to database")
	}

	fmt.Println("Successfully connected to the database!")
	return db
}
