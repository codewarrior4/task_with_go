package jobs

import (
	"log"
	"time"
	"github.com/robfig/cron/v3"
	"task/config"
	"task/models"

)

func StartVerificationCleanup() {
	c := cron.New()

	// Runs every 20 minutes
	_, err := c.AddFunc("@every 20m", cleanupExpiredVerifications)
	if err != nil {
		log.Fatal("❌ Failed to schedule verification cleanup:", err)
	}

	c.Start()
	log.Println("🔁 Started verification cleanup job with robfig/cron")
}

func cleanupExpiredVerifications() {
	log.Println("🧹 Running expired verification cleanup at", time.Now())

	db := config.DB
	result := db.Where("expires_at < ?", time.Now()).Delete(&models.Verification{})

	if result.Error != nil {
		log.Println("❌ Failed to clean expired verifications:", result.Error)
	} else {
		log.Printf("✅ Deleted %d expired verifications\n", result.RowsAffected)
	}
}
