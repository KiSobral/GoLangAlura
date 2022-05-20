package main

import (
	"GoLangAlura/ginApi/database"
	"GoLangAlura/ginApi/models"
	"GoLangAlura/ginApi/routes"
)

func main() {
	database.ConnectToDatabase()
	models.Alunos = []models.Aluno{
		{Nome: "Hugo Sobral", CPF: "00000000000", RG: "00000000"},
		{Nome: "Maria Eduarda", CPF: "11111111111", RG: "11111111"},
		{Nome: "Felipe Ferreira", CPF: "22222222222", RG: "22222222"},
	}
	routes.HandleRequest()
}
