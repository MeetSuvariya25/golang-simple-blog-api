package config

import (
	"blog-api/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Auto-migrate the Post model
	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CloseDB(db *gorm.DB) {
	// Close the database connection
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbSQL.Close()
}
