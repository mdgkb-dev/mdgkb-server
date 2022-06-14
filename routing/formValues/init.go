package formValues

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/mdgkb-server/handlers/formValues"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/documents-to-pdf/:id", h.DocumentsToPDF)
	r.GET("/:id", h.Get)
	r.PUT("/:id", h.Update)
	r.PUT("/many", h.UpdateMany)
}
