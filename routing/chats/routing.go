package chats

import (
	handler "mdgkb/mdgkb-server/handlers/chats"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(h handler.IHandler, api *gin.RouterGroup, ws *gin.RouterGroup) {
	path := "/chats"
	r := api.Group(path)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)

	ws.Group(path).GET("/connect/:chatId/:userId", h.Connect)
}
