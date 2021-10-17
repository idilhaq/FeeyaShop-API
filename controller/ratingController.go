package controller

import (
	"net/http"
	"time"

	"feeyashop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ratingInput struct {
	Rating    uint `json:"rating"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

// GetAllRating godoc
// @Summary Get all Rating.
// @Description Get a list of Rating.
// @Tags Rating
// @Produce json
// @Success 200 {object} []models.Rating
// @Router /rating [get]
func GetAllRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var ratings []models.Rating
	db.Find(&ratings)

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// CreateRating godoc
// @Summary Create New Rating.
// @Description Creating a new Rating.
// @Tags Rating
// @Param Body body ratingInput true "the body to create a new Rating"
// @Produce json
// @Success 200 {object} models.Rating
// @Router /rating [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func CreateRating(c *gin.Context) {
	// Validate input
	var input ratingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	rating := models.Rating{Rating: input.Rating, UserID: input.UserID, ProductID: input.ProductID}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&rating)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// UpdateRating godoc
// @Summary Update Rating.
// @Description Update Rating by id.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Param Body body ratingInput true "the body to update rating"
// @Success 200 {object} models.Rating
// @Router /rating/{id} [patch]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func UpdateRating(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input ratingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Rating
	updatedInput.Rating = input.Rating
	updatedInput.UpdatedAt = time.Now()

	db.Model(&rating).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// DeleteRating godoc
// @Summary Delete one Rating.
// @Description Delete a Rating by id.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Success 200 {object} map[string]boolean
// @Router /rating/{id} [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func DeleteRating(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&rating)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
