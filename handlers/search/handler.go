package search

import (
	"encoding/json"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	var item models.SearchModel
	err := json.Unmarshal([]byte(c.Query("searchModel")), &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	if item.Mode == models.SearchModeMain {
		err = h.service.MainSearch(&item)
		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
		c.JSON(http.StatusOK, item)
		return
	}
	err = h.service.SearchObjects(&item)
	if item.Mode == models.SearchModeObjects {
		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
		c.JSON(http.StatusOK, item)
		return
	}
}

func (h *Handler) SearchGroups(c *gin.Context) {
	items, err := h.service.SearchGroups()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}
