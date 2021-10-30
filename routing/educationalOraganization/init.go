package educationalOraganization

import (
	handler "mdgkb/mdgkb-server/handlers/educationalOrganization"
	"mdgkb/mdgkb-server/helpers/uploadHelper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader uploadHelper.Uploader) {
	var h = handler.CreateHandler(db, &uploader)
	r.GET("/", h.Get)
	r.PUT("", h.Update)
}