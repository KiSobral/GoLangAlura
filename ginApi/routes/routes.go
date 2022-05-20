package routes

import (
	"GoLangAlura/ginApi/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:nome", controllers.Hello)
	r.POST("/alunos", controllers.Create)
	r.GET("/alunos", controllers.ReadAll)
	r.GET("/alunos/:id", controllers.ReadOne)
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCPF)
	r.PATCH("/alunos/:id", controllers.Update)
	r.DELETE("/alunos/:id", controllers.Delete)
	r.GET("/index", controllers.ShowIndexPage)
	r.NoRoute(controllers.NotFoundRoute)
	r.Run()
}
