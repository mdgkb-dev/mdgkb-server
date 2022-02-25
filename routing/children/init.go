package children

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/children"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/", h.Create)
}
