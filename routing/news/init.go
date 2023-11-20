package news

import (
	handler "mdgkb/mdgkb-server/handlers/news"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/main", h.GetMain)
	r.GET("/submain", h.GetSubMain)
	r.GET("/relation-news", h.GetAll)
	r.GET("", h.GetAll)
	r.GET("/:slug", h.GetBySLug)
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
