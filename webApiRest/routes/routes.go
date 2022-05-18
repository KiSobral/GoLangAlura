package routes

import (
	"GoLangAlura/webApiRest/controllers"
	"GoLangAlura/webApiRest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	router.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	router.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	router.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	router.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
