package users

import (
	handler "mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/helpers"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader helpers.Uploader) {
	var h = handler.NewHandler(handler.NewRepository(db), uploader)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.GET("/get-by-email/:email", h.GetByEmail)
}
