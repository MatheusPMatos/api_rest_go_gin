package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	IdClient uint   `gorm:"foreignkey:IdClient;references:client_subdomain.ID"`
	Nome     string `json:"nome"`
	CPF      string `json:"cpf"`
	RG       string `json:"rg"`
}
