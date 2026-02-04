package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func BuscarPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	databases.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "aluno nao encontrado",
		})
		return
	}

	c.JSON(200, aluno)

}
