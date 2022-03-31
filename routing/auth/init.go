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
	r.POST("/logout", h.Logout)
	r.POST("/restore-password", h.RestorePassword)
	r.PUT("/refresh-password", h.RefreshPassword)
	r.GET("/check-uuid/:user-id/:uuid", h.CheckUUID)
	r.GET("/check-path-permissions")

	r.GET("/path-permissions", h.SavePathPermissions)
	r.PUT("/path-permissions", h.GetAllPathPermissions)
	//r.GET("/logout", handler.Logout)
}
