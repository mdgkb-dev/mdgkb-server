package news

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/news"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:slug", h.GetBySLug)
	r.GET("/month", h.GetByMonth)
	r.POST("", h.Create)
	r.PUT("/:id", h.Update)
	r.POST("/like", h.CreateLike)
	r.POST("/tag", h.AddTag)
	r.DELETE("/tag", h.RemoveTag)
	r.DELETE("/comment/:id", h.RemoveComment)
	r.PUT("/comment/:id", h.UpdateComment)
	r.POST("/comment", h.CreateComment)
	r.DELETE("/:id", h.Delete)
	r.DELETE("/like/:id", h.DeleteLike)
}
