package fileinfos

import (
	handler "mdgkb/mdgkb-server/handlers/fileinfos"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("", h.Create)
}
