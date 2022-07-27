package newsslides

import (
	handler "mdgkb/mdgkb-server/handlers/newsslides"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.PUT("/", h.UpdateAll)
}
