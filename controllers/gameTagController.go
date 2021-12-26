package controllers

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type gameTagInput struct {
	GameID uint `json:"game_id"`
	TagID  uint `json:"tag_id"`
}

// GetAllGameTag godoc
// @Summary Get all gameTag.
// @Description Get a list of GameTag.
// @Tags GameTag
// @Produce json
// @Success 200 {object} []models.GameTag
// @Router /game-tags [get]
func GetAllGameTag(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var gameTag []models.GameTag
	db.Find(&gameTag)

	c.JSON(http.StatusOK, gin.H{"data": gameTag})
}

// CreateGameTag godoc
// @Summary Create New GameTag.
// @Description Creating a new GameTag.
// @Tags GameTag
// @Param Body body gameTagInput true "the body to create a new GameTag"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.GameTag
// @Router /game-tags [post]
func CreateGameTag(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input gameTagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check Company
	var game models.Game
	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found!"})
		return
	}

	var tag models.Tag
	if err := db.Where("id = ?", input.TagID).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TagID not found!"})
		return
	}

	// Create gameTag
	gameTag := models.GameTag{
		GameID: input.GameID,
		TagID:  input.TagID,
	}
	db.Create(&gameTag)

	c.JSON(http.StatusOK, gin.H{"data": gameTag})
}

// GetGameTagById godoc
// @Summary Get GameTag.
// @Description Get a GameTag by id.
// @Tags GameTag
// @Produce json
// @Param id path string true "gameTag id"
// @Success 200 {object} models.GameTag
// @Router /game-tags/{id} [get]
func GetGameTagById(c *gin.Context) { // Get model if exist
	var gameTag models.GameTag

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&gameTag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gameTag})
}

// DeleteGameTag godoc
// @Summary Delete one gameTag.
// @Description Delete a gameTag by id.
// @Tags GameTag
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "gameTag id"
// @Success 200 {object} map[string]boolean
// @Router /game-tags/{id} [delete]
func DeleteGameTag(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var gameTag models.GameTag
	if err := db.Where("id = ?", c.Param("id")).First(&gameTag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&gameTag)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
