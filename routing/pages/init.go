package pages

import (
	handler "mdgkb/mdgkb-server/handlers/pages"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("/", h.GetAll)
	r.POST("/ftsp", h.FTSP)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)

	r.GET("/slug/:slug", h.GetBySlug)
}
