package preparations

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/preparations"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/tags", h.GetTags)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.PUT("/", h.UpdateMany)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
