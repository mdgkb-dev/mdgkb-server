package meta

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/mdgkb-server/handlers/meta"
)

// Init func
func Init(h *handler.Handler, api *gin.RouterGroup, ws *gin.RouterGroup) {
	path := "/meta"
	r := api.Group(path)
	r.GET("/count/:table", h.GetCount)
	r.GET("/schema", h.GetSchema)
	r.GET("/social", h.GetSocial)
	r.POST("/address", h.GetAddress)
	r.GET("/main", h.SearchMain)
	r.GET("/options", h.GetOptions)
	r.GET("/get-applications-counts", h.GetApplicationsCounts)

	ws.Group(path).GET("/app-counts-regular-update", h.GetWeb)
}
