package controllers

import (
	"time"

	"github.com/ARMeeru/time-capsule/models"
	"github.com/ARMeeru/time-capsule/utils"
	"github.com/gin-gonic/gin"
)

func ShowCreateCapsulePage(c *gin.Context) {
	utils.RenderTemplate(c, []string{"templates/base.html", "templates/create_capsule.html"}, nil)
}

func CreateCapsule(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	message := c.PostForm("message")
	deliverAtStr := c.PostForm("deliver_at")

	// Parse delivery time
	deliverAt, err := time.Parse("2006-01-02T15:04", deliverAtStr)
	if err != nil {
		utils.RenderTemplate(c, []string{"templates/base.html", "templates/create_capsule.html"}, gin.H{
			"Error": "Invalid delivery date format",
		})
		return
	}

	if deliverAt.Before(time.Now()) {
		utils.RenderTemplate(c, []string{"templates/base.html", "templates/create_capsule.html"}, gin.H{
			"Error": "Delivery date must be in the future",
		})
		return
	}

	capsule := models.Capsule{
		UserID:    userID,
		Message:   message,
		DeliverAt: deliverAt,
	}

	if err := utils.DB.Create(&capsule).Error; err != nil {
		utils.RenderTemplate(c, []string{"templates/base.html", "templates/create_capsule.html"}, gin.H{
			"Error": "Failed to create capsule",
		})
		return
	}

	// Render success message
	utils.RenderTemplate(c, []string{"templates/base.html", "templates/create_capsule.html"}, gin.H{
		"Success": "Capsule created successfully",
	})
}
