package controllers

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type companyInput struct {
	Name string `json:"name"`
}

// GetAllCompany godoc
// @Summary Get all Company.
// @Description Get a list of Company.
// @Tags Company
// @Produce json
// @Success 200 {object} []models.Company
// @Router /companies [get]
func GetAllCompany(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var companies []models.Company
	db.Find(&companies)

	c.JSON(http.StatusOK, gin.H{"data": companies})
}

// CreateCompany godoc
// @Summary Create New Company.
// @Description Creating a new Company.
// @Tags Company
// @Param Body body companyInput true "the body to create a new Company"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Company
// @Router /companies [post]
func CreateCompany(c *gin.Context) {
	// Validate input
	var input companyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Company
	company := models.Company{
		Name: input.Name,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&company)

	c.JSON(http.StatusOK, gin.H{"data": company})
}

// GetCompanyById godoc
// @Summary Get Company.
// @Description Get an Company by id.
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Success 200 {object} models.Company
// @Router /companies/{id} [get]
func GetCompanyById(c *gin.Context) { // Get model if exist
	var company models.Company

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": company})
}

// UpdateCompany godoc
// @Summary Update Company.
// @Description Update Company by id.
// @Tags Company
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Company id"
// @Param Body body companyInput true "the body to update company"
// @Success 200 {object} models.Company
// @Router /companies/{id} [patch]
func UpdateCompany(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var company models.Company
	if err := db.Where("id = ?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input companyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Company
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&company).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": company})
}

// DeleteCompany godoc
// @Summary Delete one Company.
// @Description Delete a Company by id.
// @Tags Company
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Company id"
// @Success 200 {object} map[string]boolean
// @Router /companies/{id} [delete]
func DeleteCompany(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var company models.Company
	if err := db.Where("id = ?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&company)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
