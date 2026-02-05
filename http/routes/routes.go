package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/api-rest-gin/internal/handler/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.FindAllStudents)
	r.POST("/students", controllers.CreatedStudent)
	r.GET("/students/:id", controllers.FindById)
	r.GET("/students/cpf/:cpf", controllers.FindByCpf)
	r.DELETE("/students/:id", controllers.DeleteById)
	r.PATCH("/students/:id", controllers.EdtingStudent)
	r.Run()
}
