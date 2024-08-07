package vacancyresponses

import (
	handler "mdgkb/mdgkb-server/handlers/vacancyresponses"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/ftsp", h.FTSP)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.GET("/email-exists/:email/:vacancyId", h.EmailExists)

	r.GET("/pdf/:id", h.PDF)
}
