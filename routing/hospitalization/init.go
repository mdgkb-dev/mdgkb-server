package hospitalization

import (
	handler "mdgkb/mdgkb-server/handlers/hospitalization"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, helper *helper.Helper) {
	var h = handler.CreateHandler(helper)
	r.GET("/", h.GetAll)
	r.GET("/pdf/:id", h.PDF)
}
