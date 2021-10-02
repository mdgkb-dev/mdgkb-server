package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

import (
	handler "mdgkb/mdgkb-server/handlers/pages"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB) {
	var h = handler.CreateHandler(db)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)

	r.GET("/slug/:slug", h.GetBySlug)
}
