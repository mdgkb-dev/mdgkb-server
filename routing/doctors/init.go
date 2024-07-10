package doctors

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/doctors"
)

// Init func
func Init(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("", h.GetAll)
	r.POST("/ftsp", h.FTSP)
	r.GET("/:slug", h.Get)
	r.GET("/division/:id", h.GetByDivisionID)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
