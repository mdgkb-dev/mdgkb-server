package search

import (
	handler "mdgkb/mdgkb-server/handlers/search"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/main", h.SearchMain)
	r.GET("/full", h.FullTextSearch)
	r.GET("/", h.Search)
	r.GET("/search-groups", h.SearchGroups)
	r.GET("/search-by-group", h.SearchGroups)
}
