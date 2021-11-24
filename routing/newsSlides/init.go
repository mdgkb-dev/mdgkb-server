package newsSlides

import (
	handler "mdgkb/mdgkb-server/handlers/newsSlides"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}