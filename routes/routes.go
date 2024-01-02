package routes

import (
	"github.com/MatheusPMatos/api-go-gin/controllers"
	"github.com/MatheusPMatos/api-go-gin/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleResquests(db *gorm.DB, r *gin.Engine) {
	controllers := controllers.NewController(repository.NewRepository(db))
	r.GET("/alunos", controllers.FindAllAlunos)
	r.GET("/:name", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaAluno)

}
