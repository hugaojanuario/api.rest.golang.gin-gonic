package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/models"
)

func DeletarPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	databases.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"massage":"delete if sucessed",
	})
}
