package controller

import (
	"net/http"

	"github.com/MaeMethas/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /users

func CreateRegistration(c *gin.Context) {

	var registrations entity.Registration
	var State entity.State
	var subject entity.Subject
	var Student entity.Student

	if err := c.ShouldBindJSON(&registrations); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", subject.ID).First(&subject); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", State.ID).First(&State); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}
	if tx := entity.DB().Where("id = ?", Student.ID).First(&Student); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}
	c.JSON(http.StatusOK, gin.H{"data": registrations})

}

// GET /registration/:id

func GetRegistration(c *gin.Context) {

	var registration entity.Registration

	id := c.Param("id")

	//ค้นหา registration ด้วย id
	if err := entity.DB().Raw("SELECT * FROM registrations WHERE id = ?", id).Scan(&registration).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": registration})

}

// List /registrations

func ListRegistration(c *gin.Context) {

	var registrations []entity.Registration

	if err := entity.DB().Preload("Student").Preload("Subject").Preload("State").Raw("SELECT * FROM registrations").Find(&registrations).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": registrations})

}

// DELETE /users/:id

func DeleteRegistration(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM registrations WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "registration not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateRegistration(c *gin.Context) {

	var registration entity.Registration

	if err := c.ShouldBindJSON(&registration); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", registration.ID).First(&registration); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "registration not found"})

		return

	}
}
