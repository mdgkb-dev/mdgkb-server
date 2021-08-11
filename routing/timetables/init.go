package timetables

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	handler "mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/helpers"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader helpers.Uploader) {
	var h = handler.NewHandler(handler.NewRepository(db), uploader)
	r.GET("/weekdays", h.GetAllWeekdays)
}
