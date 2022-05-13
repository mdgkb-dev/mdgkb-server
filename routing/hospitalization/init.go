package hospitalization

import (
	handler "mdgkb/mdgkb-server/handlers/hospitalization"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, helper *helper.Helper) {
	var h = handler.CreateHandler(db, helper)
	r.GET("/", h.GetAll)
	r.GET("/pdf/:id", h.PDF)
}
