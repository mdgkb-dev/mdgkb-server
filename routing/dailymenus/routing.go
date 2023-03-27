package dailymenus

import (
	handler "mdgkb/mdgkb-server/handlers/dailymenus"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/today-menu", h.GetTodayMenu)
	r.GET("/:id", h.Get)
	r.POST("/pdf", h.PDF)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/many", h.UpdateAll)
	r.PUT("/:id", h.Update)
}
