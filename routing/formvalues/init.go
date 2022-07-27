package formvalues

import (
	handler "mdgkb/mdgkb-server/handlers/formvalues"

	"github.com/gin-gonic/gin"
	//
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/documents-to-pdf/:id", h.DocumentsToPDF)
	r.GET("/documents-to-zip/:id", h.DocumentsToZip)
	r.GET("/:id", h.Get)
	r.PUT("/:id", h.Update)
	r.PUT("/many", h.UpdateMany)
}
