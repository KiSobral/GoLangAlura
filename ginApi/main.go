package main

import (
	"GoLangAlura/ginApi/database"
	"GoLangAlura/ginApi/routes"
)

func main() {
	database.ConnectToDatabase()

	routes.HandleRequest()
}
