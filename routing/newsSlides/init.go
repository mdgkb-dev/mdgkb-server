package newsSlides

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/mdgkb-server/handlers/newsSlides"
)

func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.PUT("/", h.UpdateAll)
}
