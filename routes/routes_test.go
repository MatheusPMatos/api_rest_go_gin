package routes

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mocks "github.com/MatheusPMatos/api-go-gin/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoutesSucess(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("A função causou um panic: %v", r)
		}
	}()

	dbMock, _, _ := sqlmock.New()
	gin.SetMode(gin.TestMode)
	r := gin.New()

	defer dbMock.Close()
	gorm, _ := mocks.OpenGorm(dbMock)
	HandleResquests(gorm, r)
	assert.True(t, true)

}
