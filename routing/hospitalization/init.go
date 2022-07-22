package hospitalization

import (
	"github.com/pro-assistance/pro-assister/helper"
	handler "mdgkb/mdgkb-server/handlers/hospitalization"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, helper *helper.Helper) {
	var h = handler.CreateHandler(helper)
	r.GET("/", h.GetAll)
	r.GET("/pdf/:id", h.PDF)
}
