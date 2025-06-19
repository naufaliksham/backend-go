package config

import (
	"backend-go/models"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		log.Println("Connected to database:", os.Getenv("DB_NAME"))
	}

	if err := DB.AutoMigrate(&models.User{}, &models.Schedule{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	} else {
		log.Println("Database migrations completed successfully")
	}

	// create user seeder
	password := "naufalbackend"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{
		Name:     "Naufal Iksham",
		Email:    "naufalbackend@gmail.com",
		Password: string(hash),
	}
	if err := DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		if err := DB.Create(&user).Error; err != nil {
			log.Fatal("Failed to create user:", err)
		} else {
			log.Println("User created successfully:", user.Name)
		}
	} else {
		log.Println("User already exists:", user.Name)
	}

}
