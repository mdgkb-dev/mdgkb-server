package dpoDocumentTypes

import (
	handler "mdgkb/mdgkb-server/handlers/dpoDocumentTypes"

	"github.com/gin-gonic/gin"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("", h.Update)
}
