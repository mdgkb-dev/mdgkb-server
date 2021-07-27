package normativeDocumentTypes

import (
	handler "mdgkb/mdgkb-server/handlers/normativeDocumentTypes"
	"mdgkb/mdgkb-server/helpers"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func Init(r *gin.RouterGroup, db *bun.DB, uploader helpers.Uploader) {
	var h = handler.NewHandler(handler.NewRepository(db), uploader)
	r.POST("/", h.Create)
}
