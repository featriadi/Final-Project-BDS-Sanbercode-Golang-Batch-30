package controllers

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type tagInput struct {
	Name string `json:"name"`
}

// GetAllTag godoc
// @Summary Get all Tag.
// @Description Get a list of Tag.
// @Tags Tag
// @Produce json
// @Success 200 {object} []models.Tag
// @Router /tags [get]
func GetAllTag(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var tags []models.Tag
	db.Find(&tags)

	c.JSON(http.StatusOK, gin.H{"data": tags})
}

// CreateTag godoc
// @Summary Create New Tag.
// @Description Creating a new Tag.
// @Tags Tag
// @Param Body body tagInput true "the body to create a new Tag"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Tag
// @Router /tags [post]
func CreateTag(c *gin.Context) {
	// Validate input
	var input tagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Tag
	tag := models.Tag{
		Name: input.Name,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&tag)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}

// GetTagById godoc
// @Summary Get Tag.
// @Description Get an Tag by id.
// @Tags Tag
// @Produce json
// @Param id path string true "Tag id"
// @Success 200 {object} models.Tag
// @Router /tags/{id} [get]
func GetTagById(c *gin.Context) { // Get model if exist
	var tag models.Tag

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tag})
}

// DeleteTag godoc
// @Summary Delete one Tag.
// @Description Delete a Tag by id.
// @Tags Tag
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Tag id"
// @Success 200 {object} map[string]boolean
// @Router /tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var tag models.Tag
	if err := db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&tag)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
