package controller

import (
	"net/http"
	"time"

	"feeyashop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentInput struct {
	Comment   string `json:"comment"`
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
}

// GetAllComment godoc
// @Summary Get all Comment.
// @Description Get a list of Comment.
// @Tags Comment
// @Produce json
// @Success 200 {object} []models.Comment
// @Router /comment [get]
func GetAllComment(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var comments []models.Comment
	db.Find(&comments)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// CreateComment godoc
// @Summary Create New Comment.
// @Description Creating a new Comment.
// @Tags Comment
// @Param Body body commentInput true "the body to create a new Comment"
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comment [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func CreateComment(c *gin.Context) {
	// Validate input
	var input commentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	comment := models.Comment{Comment: input.Comment, UserID: input.UserID, ProductID: input.ProductID}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&comment)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// UpdateComment godoc
// @Summary Update Comment.
// @Description Update Comment by id.
// @Tags Comment
// @Produce json
// @Param id path string true "Comment id"
// @Param Body body commentInput true "the body to update comment"
// @Success 200 {object} models.Comment
// @Router /comment/{id} [patch]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func UpdateComment(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var comment models.Comment
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input commentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Comment
	updatedInput.Comment = input.Comment
	updatedInput.UpdatedAt = time.Now()

	db.Model(&comment).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DeleteComment godoc
// @Summary Delete one Comment.
// @Description Delete a Comment by id.
// @Tags Comment
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} map[string]boolean
// @Router /comment/{id} [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func DeleteComment(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comment
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
