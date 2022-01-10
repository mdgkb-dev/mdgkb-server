package auth

import (
	handler "mdgkb/mdgkb-server/handlers/auth"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.POST("/refresh-token", h.RefreshToken)
	r.GET("/refresh-password", h.RefreshPassword)
	r.POST("/logout", h.Logout)
	//r.POST("/check-email", handler.CheckEmail)
	//r.GET("/logout", handler.Logout)
}
