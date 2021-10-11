package carousels

import (
	handler "mdgkb/mdgkb-server/handlers/carousels"
	"mdgkb/mdgkb-server/helpers"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader helpers.Uploader) {
	var h = handler.NewHandler(handler.NewRepository(db), uploader)
	r.GET("/", h.GetAll)
	r.GET("/key/:key", h.GetByKey)
	r.GET("/:id", h.Get)
	r.PUT("/:id", h.Update)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
}
