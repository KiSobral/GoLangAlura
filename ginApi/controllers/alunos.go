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

func ReadOne(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
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

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"Sucesso": "Aluno deletado"})
}
