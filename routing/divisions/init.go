package divisions

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/divisions"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.DELETE("/comment/:id", h.RemoveComment)
	r.PUT("/comment/:id", h.UpdateComment)
	r.POST("/comment", h.CreateComment)
}
