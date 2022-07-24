package auth

import (
	handler "mdgkb/mdgkb-server/handlers/auth"

	"github.com/gin-gonic/gin"
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
	r.POST("/check-path-permissions", h.CheckPathPermissions)

	r.GET("/path-permissions/admin", h.GetAllPathPermissionsAdmin)
	r.GET("/path-permissions", h.GetAllPathPermissions)
	r.PUT("/path-permissions", h.SavePathPermissions)
	r.GET("/path-permissions/:roleId", h.GetPathPermissionsByRoleID)
	//r.GET("/logout", handler.Logout)
}
