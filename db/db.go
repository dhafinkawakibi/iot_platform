package db

import (
	"fmt"
	"os"

	"github.com/dhafinkawakibi/iot_platform/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// InitDb ...
func InitDb(db **gorm.DB) {
	var dbErr error
	dotenvErr := godotenv.Load()

	if dotenvErr != nil {
		fmt.Println(dotenvErr)
		panic("Failed loading dotenv.")
	}

	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURL := fmt.Sprintf("%s:%s@/%s", dbUsername, dbPassword, dbName)

	*db, dbErr = gorm.Open("mysql", dbURL)

	if dbErr != nil {
		fmt.Println(dbErr)
		panic("Failed to connect to database!")
	}

	(*db).SingularTable(true)

	// Migrations
	(*db).AutoMigrate(&models.Users{})
}
