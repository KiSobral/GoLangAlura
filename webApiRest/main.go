package main

import (
	"GoLangAlura/webApiRest/models"
	"GoLangAlura/webApiRest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Maria", Historia: "A maior bióloga do Brasil"},
		{Nome: "Eduarda", Historia: "A maior artista do Brasil"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
