package controllers

import (
	"GoLangAlura/ginApi/models"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API says:": "Hello " + nome,
	})
}

func ReadAll(c *gin.Context) {
	c.JSON(200, models.Alunos)
}
