package formstatuses

import (
	handler "mdgkb/mdgkb-server/handlers/formstatuses"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/by-group/:id", h.GetAllByGroupID)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.PUT("/", h.UpdateMany)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
