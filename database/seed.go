package database

import (
	"log"
	"task/config"
	"task/models"
	"task/utils"
)

// Seed function to populate the database with sample data
func Seed() {
	// Create sample user
	hashedPassword, err := utils.HashPassword("Pa$$w0rd!")
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
	}

	user := models.User{
		Username:  "john_doe",
		Email:     "john@example.com",
		Firstname: "John",
		Lastname:  "Elijah",
		Password:  hashedPassword, // Use the hashed password
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		// User doesn't exist, create a new one
		if err := config.DB.Create(&user).Error; err != nil {
			log.Fatal("Failed to seed user: ", err)
		}
		log.Println("User seeded successfully!")
		// // Create sample tasks for the user
		// task1 := models.Task{Title: "Task 1", Description: "First task", UserID: user.ID}
		// task2 := models.Task{Title: "Task 2", Description: "Second task", UserID: user.ID}

		// // // Add tasks to the database
		// if err := config.DB.Create(&task1).Error; err != nil {
		// 	log.Fatal("Failed to seed task 1: ", err)
		// }
		// if err := config.DB.Create(&task2).Error; err != nil {
		// 	log.Fatal("Failed to seed task 2: ", err)
		// }

		// log.Println("Tasks seeded successfully!")
	} else {
		user = existingUser // Assign the existing user to user
		log.Println("User already exists, skipping seed.")
	}

}
