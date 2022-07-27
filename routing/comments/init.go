package comments

import (
	handler "mdgkb/mdgkb-server/handlers/comments"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/main", h.GetAllMain)
	r.GET("/", h.GetAll)
	r.PUT("/:id", h.UpdateOne)
	r.POST("", h.UpsertOne)
}
