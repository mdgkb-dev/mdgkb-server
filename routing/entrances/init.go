package entrances

import (
	handler "mdgkb/mdgkb-server/handlers/entrances"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
}
