package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func FindAllStudents(c *gin.Context) {
	var student []models.Student
	databases.DB.Find(&student)
	c.JSON(http.StatusOK, student)
}
