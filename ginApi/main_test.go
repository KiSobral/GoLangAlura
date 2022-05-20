package main

import (
	"GoLangAlura/ginApi/controllers"
	"GoLangAlura/ginApi/database"
	"GoLangAlura/ginApi/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func criaAlunoMock() {
	var aluno models.Aluno = models.Aluno{Nome: "Aluno Mock", CPF: "01234567890", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func deletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestHelloStatusCode(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Hello)
	req, _ := http.NewRequest("GET", "/hugo", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Os códigos deveriam ser iguais")
}

func TestHelloResponseContent(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Hello)
	req, _ := http.NewRequest("GET", "/hugo", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	responseMock := `{"API says":"Hello hugo"}`
	resposeBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, responseMock, string(resposeBody), "O corpo das mensagens deveriam ser iguais")
}

func TestGetAllAlunos(t *testing.T) {
	database.ConnectToDatabase()

	criaAlunoMock()
	defer deletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ReadAll)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Should have the same status code")
}
