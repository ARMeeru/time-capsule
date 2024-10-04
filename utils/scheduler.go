package utils

import (
	"log"
	"time"

	"github.com/ARMeeru/time-capsule/models"
)

func StartScheduler() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		DeliverCapsules()
	}
}

func DeliverCapsules() {
	var capsules []models.Capsule
	now := time.Now()

	// Find capsules due for delivery
	if err := DB.Where("deliver_at <= ? AND delivered = ?", now, false).Find(&capsules).Error; err != nil {
		log.Println("Error fetching capsules: ", err)
		return
	}

	for _, capsule := range capsules {
		var user models.User
		if err := DB.First(&user, capsule.UserID).Error; err != nil {
			log.Println("Error fetching user: ", err)
			continue
		}

		// Use the user's email as the recipient
		recipientEmail := user.Email

		// Send email
		err := SendEmail(recipientEmail, "Your Time Capsule", capsule.Message)
		if err != nil {
			log.Println("Error sending email: ", err)
			continue
		}

		// Mark capsule as delivered
		capsule.Delivered = true
		if err := DB.Save(&capsule).Error; err != nil {
			log.Println("Error updating capsule: ", err)
		} else {
			log.Printf("Capsule ID %d marked as delivered", capsule.ID)
		}

	}
}
