package jobs

import (
	"fmt"
	"log"
	"time"

	"task/config"
	"task/models"
)

func StartVerificationCleanup() {
	go func() {
		for {
			fmt.Println("🧹 [Verification Job] Running cleanup at:", time.Now().Format(time.RFC1123))

			// Delete expired verification codes
			cleanupExpiredVerifications()

			// Wait for the next run
			time.Sleep(20 * time.Minute)
		}
	}()
}

func cleanupExpiredVerifications() {
	db := config.DB

	result := db.Where("expires_at < ?", time.Now()).Delete(&models.Verification{})

	if result.Error != nil {
		log.Println("❌ Failed to delete expired verifications:", result.Error)
		return
	}

	log.Printf("✅ Deleted %d expired verification codes\n", result.RowsAffected)
}
