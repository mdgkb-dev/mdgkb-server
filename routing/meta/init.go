package meta

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/meta"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/count/:table", h.GetCount)
	r.GET("/schema", h.GetSchema)
	r.GET("/social", h.GetSocial)
	r.GET("/options", h.GetOptions)
	r.GET("/get-applications-counts", h.GetApplicationsCounts)
}
