package controllers

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewInput struct {
	UserID      uint   `json:"user_id"`
	GameID      uint   `json:"game_id"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}

// GetAllReview godoc
// @Summary Get all review.
// @Description Get a list of Review.
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Review
// @Router /reviews [get]
func GetAllReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var review []models.Review
	db.Find(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// CreateReview godoc
// @Summary Create New Review.
// @Description Creating a new Review.
// @Tags Review
// @Param Body body reviewInput true "the body to create a new Review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Review
// @Router /reviews [post]
func CreateReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input reviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check Company
	var user models.User
	if err := db.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID not found!"})
		return
	}

	var game models.Game
	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found!"})
		return
	}

	// Create Review
	review := models.Review{
		UserID:      input.UserID,
		GameID:      input.GameID,
		Rating:      input.Rating,
		Description: input.Description,
	}
	db.Create(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// GetReviewById godoc
// @Summary Get Review.
// @Description Get a Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "review id"
// @Success 200 {object} models.Review
// @Router /reviews/{id} [get]
func GetReviewById(c *gin.Context) { // Get model if exist
	var review models.Review

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateReview godoc
// @Summary Update Review.
// @Description Update review by id.
// @Tags Review
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "review id"
// @Param Body body reviewInput true "the body to update an review"
// @Success 200 {object} models.Review
// @Router /reviews/{id} [patch]
func UpdateReview(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input reviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID not found!"})
		return
	}

	var game models.Game
	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found!"})
		return
	}

	var updatedInput models.Review
	updatedInput.UserID = input.UserID
	updatedInput.GameID = input.GameID
	updatedInput.Rating = input.Rating
	updatedInput.Description = input.Description
	updatedInput.UpdatedAt = time.Now()

	db.Model(&review).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReview godoc
// @Summary Delete one Review.
// @Description Delete a review by id.
// @Tags Review
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "review id"
// @Success 200 {object} map[string]boolean
// @Router /reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
