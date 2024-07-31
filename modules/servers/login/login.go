package login

import (
	"database/sql"
	"users_v1/modules/login/handlers"
	"users_v1/modules/login/repositories"
	"users_v1/modules/login/services"

	"github.com/gin-gonic/gin"
)

func Login(router *gin.Engine, db *sql.DB) {

	r := repositories.NewRepositorie(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	router.POST("/api/login", h.LoginH)
}
