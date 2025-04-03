package database

import (
	"log"
	"log/slog"
	"task/config"
	"task/models"
)

// Migrate runs database migrations
func Migrate() {


	err := config.DB.AutoMigrate(&models.Verification{})
	slog.Info("Migrating database...")
	if err != nil {
		log.Fatal("Migration failed: ", err)
	} else {
		log.Println("Migration completed successfully!")
	}
}
