package mapnodes

import (
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NodesRequest struct {
	MapNodes models.MapNodes `json:"mapNodes"`
}

func (h *Handler) UploadMapNodes(c *gin.Context) {
	var items NodesRequest

	_, err := h.helper.HTTP.GetForm(c, &items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.service.CreateMany(items.MapNodes)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, items)
}
