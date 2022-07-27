package valuetypes

import (
	"github.com/gin-gonic/gin"
	//
	handler "mdgkb/mdgkb-server/handlers/valuetypes"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
}
