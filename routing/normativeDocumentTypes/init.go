package normativeDocumentTypes

import (
	"github.com/pro-assistance/pro-assister/helper"
	handler "mdgkb/mdgkb-server/handlers/normativeDocumentTypes"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func Init(r *gin.RouterGroup, db *bun.DB, uploader helper.Uploader) {
	var h = handler.NewHandler(handler.NewRepository(db), uploader)
	r.POST("/", h.Create)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
