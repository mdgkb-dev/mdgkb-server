package timetables

import (
	handler "mdgkb/mdgkb-server/handlers/timetables"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/weekdays", h.GetAllWeekdays)
}
