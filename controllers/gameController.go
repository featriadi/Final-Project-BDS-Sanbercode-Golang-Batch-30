package controllers

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type gameInput struct {
	Title       string `json:"title"`
	DeveloperID uint   `json:"developer_id"`
	PublisherID uint   `json:"published_id"`
	Price       int    `json:"price"`
	Year        int    `json:"year"`
}

// GetAllGame godoc
// @Summary Get all game.
// @Description Get a list of Game.
// @Tags Game
// @Produce json
// @Success 200 {object} []models.Game
// @Router /games [get]
func GetAllGame(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var game []models.Game
	db.Find(&game)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// CreateGame godoc
// @Summary Create New Game.
// @Description Creating a new Game.
// @Tags Game
// @Param Body body gameInput true "the body to create a new Game"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Game
// @Router /games [post]
func CreateGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input gameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check Company
	var company models.Company
	if err := db.Where("id = ?", input.DeveloperID).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeveloperID not found!"})
		return
	}

	if err := db.Where("id = ?", input.PublisherID).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PublisherID not found!"})
		return
	}

	// Create game
	game := models.Game{
		Title:       input.Title,
		DeveloperID: input.DeveloperID,
		PublisherID: input.PublisherID,
		Price:       input.Price,
		Year:        input.Year,
	}
	db.Create(&game)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// GetGameById godoc
// @Summary Get Game.
// @Description Get a Game by id.
// @Tags Game
// @Produce json
// @Param id path string true "game id"
// @Success 200 {object} models.Game
// @Router /games/{id} [get]
func GetGameById(c *gin.Context) { // Get model if exist
	var game models.Game

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// UpdateGame godoc
// @Summary Update Game.
// @Description Update game by id.
// @Tags Game
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "game id"
// @Param Body body gameInput true "the body to update an game"
// @Success 200 {object} models.Game
// @Router /games/{id} [patch]
func UpdateGame(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input gameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var company models.Company
	if err := db.Where("id = ?", input.DeveloperID).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeveloperID not found!"})
		return
	}

	if err := db.Where("id = ?", input.PublisherID).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PublisherID not found!"})
		return
	}

	var updatedInput models.Game
	updatedInput.Title = input.Title
	updatedInput.DeveloperID = input.DeveloperID
	updatedInput.PublisherID = input.PublisherID
	updatedInput.Price = input.Price
	updatedInput.Year = input.Year
	updatedInput.UpdatedAt = time.Now()

	db.Model(&game).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// DeleteGame godoc
// @Summary Delete one game.
// @Description Delete a game by id.
// @Tags Game
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "game id"
// @Success 200 {object} map[string]boolean
// @Router /games/{id} [delete]
func DeleteGame(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&game)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
