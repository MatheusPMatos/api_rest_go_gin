package repository

import (
	"github.com/MatheusPMatos/api-go-gin/dto"
	"github.com/MatheusPMatos/api-go-gin/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) CreateAluno(aluno models.Aluno) error {
	result := r.db.Create(&aluno)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) FindAluno() (*[]dto.AlunoDto, error) {
	var alunos []models.Aluno
	result := r.db.Find(&alunos)
	if result.Error != nil {
		return nil, result.Error
	}
	var alunosDto []dto.AlunoDto
	for i := 0; i < len(alunos); i++ {
		aluno := dto.AlunoDto{Nome: alunos[i].Nome, CPF: alunos[i].CPF, RG: alunos[i].RG, ID: alunos[i].ID}
		alunosDto = append(alunosDto, aluno)
	}
	return &alunosDto, nil
}

type Repository interface {
	FindAluno() (*[]dto.AlunoDto, error)
	CreateAluno(aluno models.Aluno) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
