package routes

import (
	"GoLangAlura/ginApi/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ReadAll)
	r.POST("/alunos", controllers.Create)
	r.GET("/alunos/:id", controllers.ReadOne)
	r.DELETE("/alunos/:id", controllers.Delete)
	r.GET("/:nome", controllers.Hello)
	r.Run()
}
