package educationalOrganizationAcademics

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/mdgkb-server/handlers/educationalOrganizationAcademics"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
}
