package educationalorganizationacademics

import (
	handler "mdgkb/mdgkb-server/handlers/educationalorganizationacademics"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
}
