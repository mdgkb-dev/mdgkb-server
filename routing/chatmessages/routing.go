package chatmessages

import (
	handler "mdgkb/mdgkb-server/handlers/chatmessages"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(h handler.IHandler, api *gin.RouterGroup) {
	r := api.Group("/chat-messages")
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
