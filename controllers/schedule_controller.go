package controllers

import (
	"backend-go/config"
	"backend-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSchedules(c *gin.Context) {
	var schedules []models.Schedule
	config.DB.Find(&schedules)
	c.JSON(http.StatusOK, schedules)
}

func CreateSchedule(c *gin.Context) {
	var input models.Schedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func UpdateSchedule(c *gin.Context) {
	var schedule models.Schedule
	if err := config.DB.First(&schedule, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	var input models.Schedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&schedule).Updates(input)
	c.JSON(http.StatusOK, schedule)
}

func DeleteSchedule(c *gin.Context) {
	var schedule models.Schedule
	if err := config.DB.First(&schedule, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	config.DB.Delete(&schedule)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
