package events

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/events"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/application", h.CreateEventApplication)
	r.GET("/main", h.GetAllForMain)
	r.GET("/:id/applications/pdf", h.EventApplicationsPDF)
}
