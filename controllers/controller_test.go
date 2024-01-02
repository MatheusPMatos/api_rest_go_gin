package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MatheusPMatos/api-go-gin/dto"
	mocks "github.com/MatheusPMatos/api-go-gin/mock"
	"github.com/MatheusPMatos/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestControllerSaudacaoSucess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controler := NewController(&mocks.Repository{})

	r := gin.New()
	r.GET("/:name", controler.Saudacao)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Matheus", nil)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

	expectBody := `{"API diz:":"E ai Matheus, tudo beleza"}`
	assert.JSONEq(t, expectBody, rr.Body.String())

}

func TestControllerCriaSucess(t *testing.T) {

	gin.SetMode(gin.TestMode)

	repoMock := mocks.Repository{}

	controler := NewController(&repoMock)

	aluno := models.Aluno{
		IdClient: 1,
		Nome:     "Matheus",
		CPF:      "12224545",
		RG:       "12464624",
	}
	repoMock.On("CreateAluno", mock.Anything).Return(nil)
	r := gin.New()
	r.POST("/aluno", controler.CriaAluno)
	rr := httptest.NewRecorder()
	json, err := json.Marshal(aluno)
	if err != nil {
		t.Error()
	}
	req := httptest.NewRequest(http.MethodPost, "/aluno", bytes.NewBufferString(string(json)))
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Result().StatusCode)
}

func TestControllerCriaFailure(t *testing.T) {

	gin.SetMode(gin.TestMode)

	repoMock := mocks.Repository{}

	controler := NewController(&repoMock)

	aluno := models.Aluno{
		IdClient: 1,
		Nome:     "Matheus",
		CPF:      "12224545",
		RG:       "12464624",
	}
	repoMock.On("CreateAluno", mock.Anything).Return(errors.New("error"))
	r := gin.New()
	r.POST("/aluno", controler.CriaAluno)
	rr := httptest.NewRecorder()
	json, err := json.Marshal(aluno)
	if err != nil {
		t.Error()
	}
	req := httptest.NewRequest(http.MethodPost, "/aluno", bytes.NewBufferString(string(json)))
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
}

func TestControllerCriaInvalidJson(t *testing.T) {

	gin.SetMode(gin.TestMode)

	repoMock := mocks.Repository{}

	controler := NewController(&repoMock)

	repoMock.On("CreateAluno", mock.Anything).Return(errors.New("error"))
	r := gin.New()
	r.POST("/aluno", controler.CriaAluno)
	rr := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/aluno", nil)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
}

func TestControllerFindSucess(t *testing.T) {

	gin.SetMode(gin.TestMode)

	repoMock := mocks.Repository{}

	controler := NewController(&repoMock)

	alunos := []dto.AlunoDto{}
	repoMock.On("FindAluno").Return(&alunos, nil)
	r := gin.New()
	r.GET("/aluno", controler.FindAllAlunos)
	rr := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/aluno", nil)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
}

func TestControllerFindFailure(t *testing.T) {

	gin.SetMode(gin.TestMode)

	repoMock := mocks.Repository{}

	controler := NewController(&repoMock)

	repoMock.On("FindAluno").Return(nil, errors.New("error"))
	r := gin.New()
	r.GET("/aluno", controler.FindAllAlunos)
	rr := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/aluno", nil)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
}
