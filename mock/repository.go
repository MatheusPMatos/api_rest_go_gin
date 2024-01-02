package mocks

import (
	"github.com/MatheusPMatos/api-go-gin/dto"
	"github.com/MatheusPMatos/api-go-gin/models"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) FindAluno() (*[]dto.AlunoDto, error) {
	args := r.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*[]dto.AlunoDto), args.Error(1)
}

func (r *Repository) CreateAluno(aluno models.Aluno) error {
	args := r.Called(aluno)
	return args.Error(0)
}
