package controller

import (
	"net/http"
	"time"

	"feeyashop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type purchaseInput struct {
	Price     uint `json:"price"`
	Amount    uint `json:"amount"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

// GetAllPurchase godoc
// @Summary Get all Purchase.
// @Description Get a list of Purchase.
// @Tags Purchase
// @Produce json
// @Success 200 {object} []models.Purchase
// @Router /purchase [get]
func GetAllPurchase(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var purchases []models.Purchase
	db.Find(&purchases)

	c.JSON(http.StatusOK, gin.H{"data": purchases})
}

// CreatePurchase godoc
// @Summary Create New Purchase.
// @Description Creating a new Purchase.
// @Tags Purchase
// @Param Body body purchaseInput true "the body to create a new Purchase"
// @Produce json
// @Success 200 {object} models.Purchase
// @Router /purchase [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func CreatePurchase(c *gin.Context) {
	// Validate input
	var input purchaseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	purchase := models.Purchase{Price: input.Price, Amount: input.Amount, UserID: input.UserID, ProductID: input.ProductID}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&purchase)

	c.JSON(http.StatusOK, gin.H{"data": purchase})
}

// UpdatePurchase godoc
// @Summary Update Purchase.
// @Description Update Purchase by id.
// @Tags Purchase
// @Produce json
// @Param id path string true "Purchase id"
// @Param Body body purchaseInput true "the body to update purchase"
// @Success 200 {object} models.Purchase
// @Router /purchase/{id} [patch]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func UpdatePurchase(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var purchase models.Purchase
	if err := db.Where("id = ?", c.Param("id")).First(&purchase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input purchaseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Purchase
	updatedInput.Price = input.Price
	updatedInput.Amount = input.Amount
	updatedInput.UpdatedAt = time.Now()

	db.Model(&purchase).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": purchase})
}

// DeletePurchase godoc
// @Summary Delete one Purchase.
// @Description Delete a Purchase by id.
// @Tags Purchase
// @Produce json
// @Param id path string true "Purchase id"
// @Success 200 {object} map[string]boolean
// @Router /purchase/{id} [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func DeletePurchase(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var purchase models.Purchase
	if err := db.Where("id = ?", c.Param("id")).First(&purchase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&purchase)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
