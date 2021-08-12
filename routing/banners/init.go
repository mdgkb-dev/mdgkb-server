package banners

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	handler "mdgkb/mdgkb-server/handlers/banners"
	"mdgkb/mdgkb-server/helpers"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader helpers.Uploader) {
	var h = handler.NewHandler(handler.NewRepository(db), uploader)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.PUT("/", h.UpdateAllOrder)
}
