package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func BuscarPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	databases.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "aluno nao encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)

}
