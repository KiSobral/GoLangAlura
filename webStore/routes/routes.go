package routes

import (
	"GoLangAlura/webStore/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
