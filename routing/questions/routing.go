package questions

import (
	handler "mdgkb/mdgkb-server/handlers/questions"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/ftsp", h.FTSP)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.PUT("/new/:id", h.ChangeNewStatus)
	r.PUT("/profile/question-answer/:user-id", h.ReadAnswers)
	r.PUT("/publish/:id", h.Publish)
	r.PUT("/many", h.UpsertMany)
}
