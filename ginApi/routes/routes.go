package routes

import (
	"GoLangAlura/ginApi/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.POST("/alunos", controllers.Create)
	r.GET("/alunos", controllers.ReadAll)
	r.GET("/alunos/:id", controllers.ReadOne)
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCPF)
	r.PATCH("/alunos/:id", controllers.Update)
	r.DELETE("/alunos/:id", controllers.Delete)
	r.Run()
}
