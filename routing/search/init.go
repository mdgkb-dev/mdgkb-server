package search

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/mdgkb-server/handlers/search"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/main", h.SearchMain)
	r.GET("/v1", h.ElasticSearch)
	r.GET("/", h.Search)
	r.GET("/search-groups", h.SearchGroups)
	r.GET("/search-by-group", h.SearchGroups)
}
