package main

import (
	"GoLangAlura/ginApi/controllers"
	"GoLangAlura/ginApi/database"
	"GoLangAlura/ginApi/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestSearchByCPF(t *testing.T) {
	database.ConnectToDatabase()
	criaAlunoMock()
	defer deletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/01234567890", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Response code should be equal")
}

func TestSearchAlunoById(t *testing.T) {
	database.ConnectToDatabase()
	criaAlunoMock()
	defer deletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.ReadOne)
	path := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno = models.Aluno{Nome: "Aluno Mock", CPF: "01234567890", RG: "123456789"}
	var aluno models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &aluno)

	assert.Equal(t, alunoMock.Nome, aluno.Nome, "Names should be equal")
	assert.Equal(t, alunoMock.CPF, aluno.CPF, "CPF should be equal")
	assert.Equal(t, alunoMock.RG, aluno.RG, "RG should be equal")
}

func TestDeleteAluno(t *testing.T) {
	database.ConnectToDatabase()
	criaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.Delete)
	r.GET("/alunos/:id", controllers.ReadOne)
	path := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status code from delete should be OK")

	newReq, _ := http.NewRequest("GET", path, nil)
	newResposta := httptest.NewRecorder()
	r.ServeHTTP(newResposta, newReq)

	assert.Equal(t, http.StatusNotFound, newResposta.Code, "Status code from search by id should be not found")
}

func TestUpdateAluno(t *testing.T) {
	database.ConnectToDatabase()
	criaAlunoMock()
	defer deletaAlunoMock()

	var alunoMock models.Aluno = models.Aluno{Nome: "Aluno Mock", CPF: "01234567890", RG: "123456789"}
	var newAlunoMock models.Aluno = models.Aluno{Nome: "Aluna", CPF: "01234567890", RG: "123456789"}
	valorJson, _ := json.Marshal(newAlunoMock)

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.Update)
	r.GET("/alunos/:id", controllers.ReadOne)
	path := "/alunos/" + strconv.Itoa(ID)

	getReq, _ := http.NewRequest("GET", path, nil)
	getResposta := httptest.NewRecorder()
	r.ServeHTTP(getResposta, getReq)

	var aluno models.Aluno
	json.Unmarshal(getResposta.Body.Bytes(), &aluno)

	assert.Equal(t, alunoMock.Nome, aluno.Nome, "The names should be equal by now")
	assert.NotEqual(t, newAlunoMock.Nome, aluno.Nome, "The names should be different by now")
	assert.Equal(t, alunoMock.CPF, aluno.CPF, "The cpf should be equal at anytime")
	assert.Equal(t, newAlunoMock.CPF, aluno.CPF, "The cpf should be equal at anytime")
	assert.Equal(t, alunoMock.RG, aluno.RG, "The rg should be equal at anytime")
	assert.Equal(t, newAlunoMock.RG, aluno.RG, "The rg should be equal at anytime")

	patchReq, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(valorJson))
	patchResoista := httptest.NewRecorder()
	r.ServeHTTP(patchResoista, patchReq)

	json.Unmarshal(patchResoista.Body.Bytes(), &aluno)

	assert.Equal(t, http.StatusOK, patchResoista.Code, "The status code should be ok")
	assert.NotEqual(t, alunoMock.Nome, aluno.Nome, "The names should be equal")
	assert.Equal(t, newAlunoMock.Nome, aluno.Nome, "The names should be different")
	assert.Equal(t, alunoMock.CPF, aluno.CPF, "The cpf should be equal at anytime")
	assert.Equal(t, newAlunoMock.CPF, aluno.CPF, "The cpf should be equal at anytime")
	assert.Equal(t, alunoMock.RG, aluno.RG, "The rg should be equal at anytime")
	assert.Equal(t, newAlunoMock.RG, aluno.RG, "The rg should be equal at anytime")
}
