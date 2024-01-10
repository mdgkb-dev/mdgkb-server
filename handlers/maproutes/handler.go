package maproutes

import (
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMapRoute(c *gin.Context) {
	item, err := h.service.GetMapRoute(c.Param("start-node-id"), c.Param("end-node-id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

type NodesRequest struct {
	MapNodes models.MapNodes `json:"mapNodes"`
}
