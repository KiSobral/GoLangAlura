package controllers

import (
	"GoLangAlura/webStore/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", produtos)
}
