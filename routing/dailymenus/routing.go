package dailymenus

import (
	handler "mdgkb/mdgkb-server/handlers/dailymenus"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(api *gin.RouterGroup, ws *gin.RouterGroup, h handler.IHandler) {
	path := "/daily-menus"
	r := api.Group(path)
	r.GET("/", h.GetAll)
	r.GET("/today-menu", h.GetTodayMenu)
	r.GET("/:id", h.Get)
	r.POST("/pdf", h.PDF)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/many", h.UpdateAll)
	r.PUT("/:id", h.Update)
	ws.Group(path).GET("/regular-update", h.GetWeb)
}
