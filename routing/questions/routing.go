package vacancies

import (
	handler "mdgkb/mdgkb-server/handlers/questions"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.PUT("/new/:id", h.ChangeNewStatus)
	r.PUT("/read-answers/:user-id", h.ReadAnswers)
	r.PUT("/publish/:id", h.Publish)
	r.PUT("/many", h.UpsertMany)
}
