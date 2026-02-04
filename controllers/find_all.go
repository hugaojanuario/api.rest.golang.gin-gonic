package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func ExibeTodosOsAlunos(c *gin.Context) {
	var alunos [] models.Aluno
	databases.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}
