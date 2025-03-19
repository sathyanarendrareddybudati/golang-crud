package routes

import (
	"golang-crud/initializers"
	"golang-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSchool(c *gin.Context) {
	var school models.Schools
	if err := c.ShouldBindJSON(&school); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&school).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create school"})
		return
	}

	c.JSON(http.StatusCreated, school)
}

func GetSchools(c *gin.Context) {
	var schools []models.Schools
	if err := initializers.DB.Find(&schools).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schools"})
		return
	}

	c.JSON(http.StatusOK, schools)
}

func GetSchoolByID(c *gin.Context) {
	var school models.Schools
	id := c.Param("id")

	if err := initializers.DB.First(&school, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}

	c.JSON(http.StatusOK, school)
}

func UpdateSchool(c *gin.Context) {
	var school models.Schools
	id := c.Param("id")

	if err := initializers.DB.First(&school, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}

	if err := c.ShouldBindJSON(&school); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Save(&school).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update school"})
		return
	}

	c.JSON(http.StatusOK, school)
}

func DeleteSchool(c *gin.Context) {
	var school models.Schools
	id := c.Param("id")

	if err := initializers.DB.First(&school, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}

	if err := initializers.DB.Delete(&school).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete school"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "School deleted successfully"})
}
