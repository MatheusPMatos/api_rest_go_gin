package controllers

import (
	"net/http"

	"github.com/MatheusPMatos/api-go-gin/models"
	"github.com/MatheusPMatos/api-go-gin/repository"
	"github.com/gin-gonic/gin"
)

type controller struct {
	rp repository.Repository
}

type Controller interface {
	FindAllAlunos(*gin.Context)
	Saudacao(*gin.Context)
	CriaAluno(*gin.Context)
}

func NewController(repo repository.Repository) Controller {
	return &controller{rp: repo}
}

func (co *controller) Saudacao(c *gin.Context) {
	nome := c.Params.ByName("name")
	c.JSON(200, gin.H{"API diz:": "E ai " + nome + ", tudo beleza"})
}

func (co *controller) FindAllAlunos(c *gin.Context) {
	alunos, err := co.rp.FindAluno()
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, alunos)
}

func (co *controller) CriaAluno(c *gin.Context) {

	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err := co.rp.CreateAluno(aluno)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusCreated, nil)
}
