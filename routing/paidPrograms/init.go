package paidPrograms

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/paidPrograms"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/:id", h.Get)
	r.PUT("/:id", h.Update)
}
