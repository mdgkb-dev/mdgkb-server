package maproutes

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/maproutes"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/:start-node-id/:end-node-id", h.GetMapRoute)
}
