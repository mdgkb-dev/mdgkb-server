package search

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/mdgkb-server/handlers/search"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.Search)
	r.GET("/search-groups", h.SearchGroups)
}