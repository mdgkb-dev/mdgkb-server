package certificates

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/certifiactes"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.PUT("", h.UpdateMany)
}
