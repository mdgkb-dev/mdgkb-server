package paidprograms

import (
	handler "mdgkb/mdgkb-server/handlers/paidprograms"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/:id", h.Get)
	r.PUT("/:id", h.Update)
}
