package postgraduateCourses

import (
	handler "mdgkb/mdgkb-server/handlers/postgraduateCourses"

	"github.com/gin-gonic/gin"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/get", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
