package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mocks "github.com/MatheusPMatos/api-go-gin/mock"
	"github.com/MatheusPMatos/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

func TestAlunoCreateSuccess(t *testing.T) {
	dbMock, mock, _ := sqlmock.New()
	defer dbMock.Close()
	gorm, _ := mocks.OpenGorm(dbMock)
	repo := NewRepository(gorm)

	aluno := models.Aluno{
		IdClient: 1,
		Nome:     "Matheus",
		CPF:      "12224545",
		RG:       "12464624",
	}

	rowsid := mock.NewRows(
		[]string{"ID"}).
		AddRow(1)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), aluno.IdClient, aluno.Nome, aluno.CPF, aluno.RG).
		WillReturnRows(rowsid)
	mock.ExpectCommit()
	mock.ExpectClose()

	err := repo.CreateAluno(aluno)

	assert.Nil(t, err)
}
func TestAlunoCreateFailure(t *testing.T) {
	dbMock, mock, _ := sqlmock.New()
	defer dbMock.Close()
	gorm, _ := mocks.OpenGorm(dbMock)
	repo := NewRepository(gorm)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").
		WithArgs().
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()

	err := repo.CreateAluno(models.Aluno{
		Nome: "Matheus",
		CPF:  "12224545",
		RG:   "12464624",
	})

	assert.NotNil(t, err)
}

func TestAlunoFindSuccess(t *testing.T) {
	dbMock, mock, _ := sqlmock.New()
	defer dbMock.Close()
	gorm, _ := mocks.OpenGorm(dbMock)
	repo := NewRepository(gorm)

	aluno := models.Aluno{
		IdClient: 1,
		Nome:     "Matheus",
		CPF:      "12224545",
		RG:       "12464624",
	}
	aluno2 := models.Aluno{
		IdClient: 1,
		Nome:     "Julia",
		CPF:      "12224545",
		RG:       "12464624",
	}
	rows := mock.NewRows(
		[]string{"id", "client_id", "nome", "cpf", "rg"}).
		AddRow(1, aluno.IdClient, aluno.Nome, aluno.CPF, aluno.RG).
		AddRow(2, aluno2.IdClient, aluno2.Nome, aluno2.CPF, aluno2.RG)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectClose()

	alunos, err := repo.FindAluno()

	assert.Nil(t, err)
	assert.Equal(t, len(*alunos), 2)
}
func TestAlunoFindFailure(t *testing.T) {
	dbMock, mock, _ := sqlmock.New()
	defer dbMock.Close()
	gorm, _ := mocks.OpenGorm(dbMock)
	repo := NewRepository(gorm)

	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)
	mock.ExpectBegin()
	mock.ExpectClose()

	alunos, err := repo.FindAluno()

	assert.Nil(t, alunos)
	assert.NotNil(t, err)
}

func TestNewAlunoRepository(t *testing.T) {
	dbMock, _, _ := sqlmock.New()

	defer dbMock.Close()
	gorm, _ := mocks.OpenGorm(dbMock)
	repo := NewRepository(gorm)

	assert.NotNil(t, repo)
}
