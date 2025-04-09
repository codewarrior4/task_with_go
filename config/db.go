package config

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "task/config"  // Assuming your GetEnv function is located here
)

// Global DB instance
var DB *gorm.DB

// ConnectDB connects to the MySQL database using credentials from the environment
func ConnectDB() {
	// Fetch database credentials from environment using GetEnv function
	dbUser := GetEnv("DB_USER", "root")       // Default value "root"
	dbPassword := GetEnv("DB_PASS", "a")   // Default empty string
	dbHost := GetEnv("DB_HOST", "127.0.0.1") // Default value "127.0.0.1"
	dbPort := GetEnv("DB_PORT", "3306")      // Default value "3306"
	dbName := GetEnv("DB_NAME", "task")      // Default value "task"

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully!")
}
