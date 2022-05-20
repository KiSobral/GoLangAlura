package routes

import (
	"GoLangAlura/ginApi/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ReadAll)
	r.GET("/:nome", controllers.Hello)
	r.Run()
}
