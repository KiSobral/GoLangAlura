package controllers

import (
	"GoLangAlura/ginApi/database"
	"GoLangAlura/ginApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	nome := c.Param("nome")
	c.JSON(http.StatusOK, gin.H{
		"API says": "Hello " + nome})
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

	if err := c.ShouldBindJSON(&aluno); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosDeAlunos(&aluno); err != nil {
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

func Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosDeAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func SearchByCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	var aluno models.Aluno
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
