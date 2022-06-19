package residencyApplications

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/mdgkb-server/handlers/residencyApplications"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.GET("/email-exists/:email/:courseId", h.EmailExists)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.PUT("/many", h.UpsertMany)
}
