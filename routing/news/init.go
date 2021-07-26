package news

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	handler "mdgkb/mdgkb-server/handlers/news"
	"mdgkb/mdgkb-server/helpers"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader helpers.Uploader) {
	var h = handler.NewHandler(handler.NewRepository(db), uploader)
	r.GET("/", h.GetAll)
	r.GET("/:slug", h.GetBySLug)
	r.POST("", h.Create)
	r.PUT("/:id", h.Update)
	r.POST("/like", h.CreateLike)
	r.POST("/tag", h.AddTag)
	r.DELETE("/tag", h.RemoveTag)
	r.POST("/comment", h.CreateComment)
	r.DELETE("/:id", h.Delete)
	r.DELETE("/like/:id", h.DeleteLike)
	r.DELETE("/comment/:id", h.DeleteComment)
	r.PUT("/:id/status", h.UpdateStatus)
}
