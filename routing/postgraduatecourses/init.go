package postgraduatecourses

import (
	handler "mdgkb/mdgkb-server/handlers/postgraduatecourses"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("", h.GetAll)
	r.GET("/get", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.PUT("/many", h.UpsertMany)
}
