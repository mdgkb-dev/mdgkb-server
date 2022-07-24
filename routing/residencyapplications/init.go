package residencyapplications

import (
	handler "mdgkb/mdgkb-server/handlers/residencyapplications"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.GET("/email-exists/:email/:courseId", h.EmailExists)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/form/:id", h.UpdateWithForm)
	r.PUT("/:id", h.Update)
	r.PUT("/many", h.UpsertMany)
	r.POST("/fill-application-template", h.FillApplicationTemplate)
}
