package divisions

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	handler "mdgkb/mdgkb-server/handlers/divisions"
	"mdgkb/mdgkb-server/helpers/uploadHelper"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader uploadHelper.Uploader) {
	var h = handler.CreateHandler(db, &uploader)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.DELETE("/comment/:id", h.RemoveComment)
	r.PUT("/comment/:id", h.UpdateComment)
	r.POST("/comment", h.CreateComment)
}
