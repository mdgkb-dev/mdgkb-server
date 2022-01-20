package users

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/users"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.GET("/get-by-email/:email", h.GetByEmail)
	r.PUT("/:id", h.Update)
	r.POST("/add-to-user/:domain", h.AddToUser)
	r.DELETE("/remove-from-user/:domain/:id", h.RemoveFromUser)
}
