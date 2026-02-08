package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func CreatedStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	if err := models.ValidationStudent(&student); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	databases.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}
