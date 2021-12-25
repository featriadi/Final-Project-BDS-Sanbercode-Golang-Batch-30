package controllers

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllGame godoc
// @Summary Get all game.
// @Description Get a list of Game.
// @Tags Game
// @Produce json
// @Success 200 {object} []models.Game
// @Router /game [get]
func GetAllGame(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var game []models.Game
	db.Find(&game)

	c.JSON(http.StatusOK, gin.H{"data": game})
}
