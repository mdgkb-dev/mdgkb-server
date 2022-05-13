package document_types

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/documentTypes"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/for/tables-names", h.GetDocumentsTypesForTablesNames)

	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)

}
