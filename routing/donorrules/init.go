package donorrules

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/donorrules"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.PUT("", h.UpdateMany)
	r.POST("/add-to-user", h.AddToUser)
	r.DELETE("/delete-from-user/:donor-rule-id", h.DeleteFromUser)
}
