package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.POST("alunos", controllers.CriarNovoAluno)
	r.GET("/alunos/:id", controllers.BuscarPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarPorCpf)
	r.DELETE("alunos/:id", controllers.DeletarPorId)
	r.PATCH("alunos/:id", controllers.EditarAluno)
	r.Run()
}
