package appointments

import (
	handler "mdgkb/mdgkb-server/handlers/appointments"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.GET("/init", h.Init)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
