package routes

import (
	"GoLangAlura/webApiRest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	router.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", router))
}
