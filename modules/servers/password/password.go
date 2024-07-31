package password

import (
	"database/sql"
	"users_v1/modules/password/handlers"
	"users_v1/modules/password/repositories"
	"users_v1/modules/password/services"

	"github.com/gin-gonic/gin"
)

func Password(router *gin.Engine, db *sql.DB) {

	r := repositories.NewRepositorie(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	router.PATCH("/api/newPassword", h.InitPasswordH)
	router.PATCH("/api/changePassword", h.ChangePasswordHandler)
}
