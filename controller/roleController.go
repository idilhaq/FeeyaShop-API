package controller

import (
	"net/http"
	"time"

	"feeyashop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type roleInput struct {
	Role string `json:"role"`
}

// GetAllRole godoc
// @Summary Get all Role.
// @Description Get a list of Role.
// @Tags Role
// @Produce json
// @Success 200 {object} []models.Role
// @Router /role [get]
func GetAllRole(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var roles []models.Role
	db.Find(&roles)

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// CreateRole godoc
// @Summary Create New Role.
// @Description Creating a new Role.
// @Tags Role
// @Param Body body roleInput true "the body to create a new Role"
// @Produce json
// @Success 200 {object} models.Role
// @Router /role [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func CreateRole(c *gin.Context) {
	// Validate input
	var input roleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	role := models.Role{Role: input.Role}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&role)

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// UpdateRole godoc
// @Summary Update Role.
// @Description Update Role by id.
// @Tags Role
// @Produce json
// @Param id path string true "Role id"
// @Param Body body roleInput true "the body to update role"
// @Success 200 {object} models.Role
// @Router /role/{id} [patch]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func UpdateRole(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var role models.Role
	if err := db.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input roleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Role
	updatedInput.Role = input.Role
	updatedInput.UpdatedAt = time.Now()

	db.Model(&role).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// DeleteRole godoc
// @Summary Delete one Role.
// @Description Delete a Role by id.
// @Tags Role
// @Produce json
// @Param id path string true "Role id"
// @Success 200 {object} map[string]boolean
// @Router /role/{id} [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer )
func DeleteRole(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var role models.Role
	if err := db.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&role)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
