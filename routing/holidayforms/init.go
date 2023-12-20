package holidayforms

import (
	handler "mdgkb/mdgkb-server/handlers/holidayforms"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("", h.Create)
}
