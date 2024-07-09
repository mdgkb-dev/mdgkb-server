package divisions

import (
	handler "mdgkb/mdgkb-server/handlers/divisions"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("ftsp", h.FTSP)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.DELETE("/comment/:id", h.RemoveComment)
	r.PUT("/comment/:id", h.UpdateComment)
	r.POST("/comment", h.CreateComment)
}
