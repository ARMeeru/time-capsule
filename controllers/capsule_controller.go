package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ARMeeru/time-capsule/models"
	"github.com/ARMeeru/time-capsule/utils"
)

func CreateCapsule(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var input struct {
		Message   string `json:"message" binding:"required"`
		DeliverAt string `json:"deliver_at" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse delivery time
	deliverAt, err := time.Parse(time.RFC3339, input.DeliverAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid delivery date format"})
		return
	}

	if deliverAt.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delivery date must be in the future"})
		return
	}

	capsule := models.Capsule{
		UserID:    userID,
		Message:   input.Message,
		DeliverAt: deliverAt,
	}

	if err := utils.DB.Create(&capsule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create capsule"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Capsule created successfully"})
}
