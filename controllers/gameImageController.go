package controllers

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type gameImageInput struct {
	GameID   uint   `json:"game_id"`
	ImageUrl string `json:"image_url"`
}

// GetAllGameImage godoc
// @Summary Get all gameImage.
// @Description Get a list of GameImage.
// @Tags GameImage
// @Produce json
// @Success 200 {object} []models.GameImage
// @Router /game-images [get]
func GetAllGameImage(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var gameImage []models.GameImage
	db.Find(&gameImage)

	c.JSON(http.StatusOK, gin.H{"data": gameImage})
}

// CreateGameImage godoc
// @Summary Create New GameImage.
// @Description Creating a new GameImage.
// @Tags GameImage
// @Param Body body gameImageInput true "the body to create a new GameImage"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.GameImage
// @Router /game-images [post]
func CreateGameImage(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input gameImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check Company
	var company models.Company
	if err := db.Where("id = ?", input.GameID).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeveloperID not found!"})
		return
	}

	// Create gameImage
	gameImage := models.GameImage{
		GameID:   input.GameID,
		ImageUrl: input.ImageUrl,
	}
	db.Create(&gameImage)

	c.JSON(http.StatusOK, gin.H{"data": gameImage})
}

// GetGameImageById godoc
// @Summary Get GameImage.
// @Description Get a GameImage by id.
// @Tags GameImage
// @Produce json
// @Param id path string true "gameImage id"
// @Success 200 {object} models.GameImage
// @Router /game-images/{id} [get]
func GetGameImageById(c *gin.Context) { // Get model if exist
	var gameImage models.GameImage

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&gameImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gameImage})
}

// DeleteGameImage godoc
// @Summary Delete one gameImage.
// @Description Delete a gameImage by id.
// @Tags GameImage
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "gameImage id"
// @Success 200 {object} map[string]boolean
// @Router /game-images/{id} [delete]
func DeleteGameImage(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var gameImage models.GameImage
	if err := db.Where("id = ?", c.Param("id")).First(&gameImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&gameImage)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
