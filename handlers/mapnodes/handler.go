package mapnodes

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type NodesRequest struct {
// 	MapNodes models.MapNodes `json:"mapNodes"`
// }

func (h *Handler) UploadMapNodes(c *gin.Context) {
	var items models.MapNodes

	_, err := h.helper.HTTP.GetForm(c, &items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.service.CreateMany(items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	fmt.Println(err, "dddd")
	c.JSON(http.StatusOK, items)
}
