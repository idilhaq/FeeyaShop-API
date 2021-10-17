package controller

import (
	"net/http"

	"feeyashop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type likeInput struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

// GetAllLike godoc
// @Summary Get all Like.
// @Description Get a list of Like.
// @Tags Like
// @Produce json
// @Success 200 {object} []models.Like
// @Router /like [get]
func GetAllLike(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var likes []models.Like
	db.Find(&likes)

	c.JSON(http.StatusOK, gin.H{"data": likes})
}

// CreateLike godoc
// @Summary Create New Like.
// @Description Creating a new Like.
// @Tags Like
// @Param Body body likeInput true "the body to create a new Like"
// @Produce json
// @Success 200 {object} models.Like
// @Router /like [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func CreateLike(c *gin.Context) {
	// Validate input
	var input likeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	like := models.Like{UserID: input.UserID, ProductID: input.ProductID}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&like)

	c.JSON(http.StatusOK, gin.H{"data": like})
}

// DeleteLike godoc
// @Summary Delete one Like.
// @Description Delete a Like by id.
// @Tags Like
// @Produce json
// @Param id path string true "Like id"
// @Success 200 {object} map[string]boolean
// @Router /like/{id} [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func DeleteLike(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var like models.Like
	if err := db.Where("id = ?", c.Param("id")).First(&like).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&like)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
