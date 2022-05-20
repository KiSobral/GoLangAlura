package controllers

import (
	"GoLangAlura/ginApi/database"
	"GoLangAlura/ginApi/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API says:": "Hello " + nome,
	})
}

func ReadAll(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Create(c *gin.Context) {
	var aluno models.Aluno

	fmt.Println("AAAAAAAAAAAAAAAA")

	if err := c.ShouldBindJSON(&aluno); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
