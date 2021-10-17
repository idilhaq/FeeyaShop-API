package controller

import (
	"net/http"
	"time"

	"feeyashop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoryInput struct {
	Name string `json:"name"`
}

// GetAllCategory godoc
// @Summary Get all Category.
// @Description Get a list of Category.
// @Tags Category
// @Produce json
// @Success 200 {object} []models.Category
// @Router /category [get]
func GetAllCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Category
	db.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GetProductByCategoryId godoc
// @Summary Get Products.
// @Description Get all Products by CategoryId.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} []models.Product
// @Router /category/{id} [get]
func GetProductsByCategoryId(c *gin.Context) { // Get model if exist
	var categories []models.Category

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("category_id = ?", c.Param("id")).Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// CreateCategory godoc
// @Summary Create New Category.
// @Description Creating a new Category.
// @Tags Category
// @Param Body body categoryInput true "the body to create a new Category"
// @Produce json
// @Success 200 {object} models.Category
// @Router /category [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func CreateCategory(c *gin.Context) {
	// Validate input
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	category := models.Category{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// UpdateCategory godoc
// @Summary Update Category.
// @Description Update Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Param Body body categoryInput true "the body to update category"
// @Success 200 {object} models.Category
// @Router /category/{id} [patch]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func UpdateCategory(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Category
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&category).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory godoc
// @Summary Delete one Category.
// @Description Delete a Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} map[string]boolean
// @Router /category/{id} [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func DeleteCategory(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
