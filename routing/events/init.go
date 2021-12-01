package events
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/events"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/application", h.CreateEventApplication)
	r.GET("/:id/applications/pdf", h.EventApplicationsPDF)
}
