package candidateExams

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/mdgkb-server/handlers/candidateExams"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
