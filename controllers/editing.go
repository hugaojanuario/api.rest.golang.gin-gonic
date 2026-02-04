package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	databases.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	databases.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
	
}
