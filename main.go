package main

import (
	"log"

	"github.com/MatheusPMatos/api-go-gin/database"
	"github.com/MatheusPMatos/api-go-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.ConectaComBancodeDados()

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	r := gin.Default()
	routes.HandleResquests(db, r)
	r.Run()

}
