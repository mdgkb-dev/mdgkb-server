package dataexport

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/dataexport"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("export", h.Export)
	r.GET("data", h.Data)
}
