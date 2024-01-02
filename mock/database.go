package mocks

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenGorm(db *sql.DB) (*gorm.DB, error) {
	DB, err := gorm.Open(

		postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})

	return DB, err
}
