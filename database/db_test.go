package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConectaComBancodeDadosSucess(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("A função causou um panic: %v", r)
		}
	}()
	db, err := ConectaComBancodeDados()

	assert.NotNil(t, db)
	assert.Nil(t, err)

}
