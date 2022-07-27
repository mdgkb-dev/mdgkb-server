package children

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/children"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/", h.Create)
}
