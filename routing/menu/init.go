package menu

import (
	handler "mdgkb/mdgkb-server/handlers/menu"
	"mdgkb/mdgkb-server/helpers/uploadHelper"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader uploadHelper.Uploader) {
	var h = handler.CreateHandler(db, &uploader)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
