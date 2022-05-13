package buildings

import (
	handler "mdgkb/mdgkb-server/handlers/buildings"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB) {
	var h = handler.NewHandler(handler.NewRepository(db))
	r.GET("/", h.GetAll)
	r.GET("/floor/:id", h.GetByFloorId)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.Update)
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
}
