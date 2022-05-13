package donorRules

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/donorRules"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.PUT("", h.UpdateMany)
	r.POST("/add-to-user", h.AddToUser)
	r.DELETE("/delete-from-user/:donor-rule-id", h.DeleteFromUser)
}
