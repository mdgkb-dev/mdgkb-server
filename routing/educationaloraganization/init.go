package educationaloraganization

import (
	handler "mdgkb/mdgkb-server/handlers/educationalorganization"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.Get)
	r.PUT("", h.Update)
}
