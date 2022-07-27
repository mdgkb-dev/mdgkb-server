package buildings

import (
	handler "mdgkb/mdgkb-server/handlers/buildings"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/floor/:id", h.GetByFloorID)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
}
