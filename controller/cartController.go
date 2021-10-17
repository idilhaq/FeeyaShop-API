package controller

import (
	"net/http"
	"time"

	"feeyashop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type cartInput struct {
	Amount    uint `json:"amount"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

// GetAllCart godoc
// @Summary Get all Cart.
// @Description Get a list of Cart.
// @Tags Cart
// @Produce json
// @Success 200 {object} []models.Cart
// @Router /cart [get]
func GetAllCart(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var carts []models.Cart
	db.Find(&carts)

	c.JSON(http.StatusOK, gin.H{"data": carts})
}

// CreateCart godoc
// @Summary Create New Cart.
// @Description Creating a new Cart.
// @Tags Cart
// @Param Body body cartInput true "the body to create a new Cart"
// @Produce json
// @Success 200 {object} models.Cart
// @Router /cart [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func CreateCart(c *gin.Context) {
	// Validate input
	var input cartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	cart := models.Cart{Amount: input.Amount, UserID: input.UserID, ProductID: input.ProductID}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&cart)

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// UpdateCart godoc
// @Summary Update Cart.
// @Description Update Cart by id.
// @Tags Cart
// @Produce json
// @Param id path string true "Cart id"
// @Param Body body cartInput true "the body to update cart"
// @Success 200 {object} models.Cart
// @Router /cart/{id} [patch]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func UpdateCart(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var cart models.Cart
	if err := db.Where("id = ?", c.Param("id")).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input cartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Cart
	updatedInput.Amount = input.Amount
	updatedInput.UpdatedAt = time.Now()

	db.Model(&cart).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// DeleteCart godoc
// @Summary Delete one Cart.
// @Description Delete a Cart by id.
// @Tags Cart
// @Produce json
// @Param id path string true "Cart id"
// @Success 200 {object} map[string]boolean
// @Router /cart/{id} [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func DeleteCart(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var cart models.Cart
	if err := db.Where("id = ?", c.Param("id")).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&cart)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
