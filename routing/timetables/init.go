package timetables

import (
	handler "mdgkb/mdgkb-server/handlers/timetables"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB) {
	var h = handler.CreateHandler(db)
	r.GET("/weekdays", h.GetAllWeekdays)
}
