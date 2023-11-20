package search

import (
	"encoding/json"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	var item models.SearchModel
	item.Key = c.Query("key")
	item.Query = c.Query("query")
	err := h.service.SearchObjects(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) FullTextSearch(c *gin.Context) {
	var item models.SearchModel
	err := json.Unmarshal([]byte(c.Query("searchModel")), &item)

	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Search(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) SearchMain(c *gin.Context) {
	var item models.SearchModel
	err := json.Unmarshal([]byte(c.Query("searchModel")), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.SearchMain(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) SearchGroups(c *gin.Context) {
	items, err := h.service.SearchGroups()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}
