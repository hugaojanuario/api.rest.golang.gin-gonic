package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func FindById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	databases.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "aluno nao encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, student)

}
