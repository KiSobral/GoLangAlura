package main

import (
	"GoLangAlura/webApiRest/models"
	"GoLangAlura/webApiRest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Maria", Historia: "A maior bióloga do Brasil"},
		{Id: 2, Nome: "Eduarda", Historia: "A maior artista do Brasil"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
