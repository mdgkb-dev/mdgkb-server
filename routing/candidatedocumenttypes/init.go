package candidatedocumenttypes

import (
	handler "mdgkb/mdgkb-server/handlers/candidatedocumenttypes"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("", h.Update)
}
