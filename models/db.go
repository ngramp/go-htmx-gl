package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	// Initialize the models connection
	//dsn := "host=localhost user=your_username password=your_password dbname=your_db_name port=5432 sslmode=disable"
	//models, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("Failed to connect to models!")
	//}
	// Connect to the SQLite models
	database, err := gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = database.AutoMigrate(&User{}, &Article{})
	if err != nil {
		log.Fatal("Failed to auto-migrate models: ", err)
	}
	fmt.Println("Connected to database and created tables successfully")
	DB = database
	createUsers()
}

func createUsers() {
	// Create test user records
	users := []User{
		{Username: "admin", Password: "admin", Role: "admin"},
		{Username: "editor", Password: "editor", Role: "editor"},
	}

	for _, user := range users {
		result := DB.Create(&user)
		if result.Error != nil {
			log.Printf("Failed to create user %v: %v\n", user.Username, result.Error)
		} else {
			fmt.Printf("Created user: %v\n", user.Username)
		}
	}
}
