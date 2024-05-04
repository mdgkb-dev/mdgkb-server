package formpatterns

import (
	handler "mdgkb/mdgkb-server/handlers/formpatterns"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/ftsp", h.FTSP)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
