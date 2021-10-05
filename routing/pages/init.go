package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

import (
	handler "mdgkb/mdgkb-server/handlers/pages"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader uploadHelper.Uploader) {
	var h = handler.CreateHandler(db,&uploader)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)

	r.GET("/slug/:slug", h.GetBySlug)
}
