package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func FindByCpf(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	databases.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "student nao encontrado",
		})
		return
	}

	c.JSON(200, student)

}
