package users

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/users"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.GET("/get-by-email/:email", h.GetByEmail)
	r.PUT("/:id", h.Update)
	r.POST("", h.Create)
	r.POST("/add-to-user/:domain", h.AddToUser)
	r.DELETE("/remove-from-user/:domain/:id", h.RemoveFromUser)
}
