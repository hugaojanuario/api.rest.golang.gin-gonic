package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func EdtingStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	databases.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	databases.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)

}
