package controllers

import "github.com/gin-gonic/gin"


func PegaParametro(c *gin.Context){
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"api diz: ":"eae " + nome,
	})
}