package controller

import (
	"net/http"
	"time"

	"feeyashop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	CategoryID  uint   `json:"category_id"`
}

// GetAllProduct godoc
// @Summary Get all Product.
// @Description Get a list of Product.
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @Router /product [get]
func GetAllProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct godoc
// @Summary Create New Product.
// @Description Creating a new Product.
// @Tags Product
// @Param Body body productInput true "the body to create a new Product"
// @Produce json
// @Success 200 {object} models.Product
// @Router /product [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func CreateProduct(c *gin.Context) {
	// Validate input
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Product
	product := models.Product{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct godoc
// @Summary Update Product.
// @Description Update Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Param Body body productInput true "the body to update product"
// @Success 200 {object} models.Product
// @Router /product/{id} [patch]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func UpdateProduct(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Product
	updatedInput.Name = input.Name
	updatedInput.Description = input.Description
	updatedInput.Price = input.Price
	updatedInput.CategoryID = input.CategoryID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&product).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct godoc
// @Summary Delete one Product.
// @Description Delete a Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} map[string]boolean
// @Router /product/{id} [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
