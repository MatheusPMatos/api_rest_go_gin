package database

import (
	"log"

	"github.com/MatheusPMatos/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaComBancodeDados() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
		return nil, err
	}

	DB.AutoMigrate(&models.Aluno{})
	return DB, nil

}
