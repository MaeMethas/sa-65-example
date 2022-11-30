package controller

import (
	"github.com/MaeMethas/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /subjects
func GetSubject(c *gin.Context) {
	var subject entity.Subject
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM subjects WHERE id = ?", id).Scan(&subject).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": subject})
}

// List /subjects
func ListSubject(c *gin.Context) {
	var subjects []entity.Subject
	if err := entity.DB().Raw("SELECT * FROM subjects").Scan(&subjects).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subjects})
}

// DELETE /subjects/:id
func DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FORM subjects WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /subjects
func UpdateSubject(c *gin.Context) {
	var subject entity.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", subject.ID).First(&subject); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject not found"})
		return
	}

	if err := entity.DB().Save(&subject).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subject})
}
