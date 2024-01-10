package mapnodes

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/mapnodes"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/upload", h.UploadMapNodes)
}
