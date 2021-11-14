package comments

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/mdgkb-server/handlers/comments"
)

func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.PUT("/:id", h.UpdateOne)
}
