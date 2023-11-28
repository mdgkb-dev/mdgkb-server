package mapnodes

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NodesRequest struct {
	MapNodes models.MapNodes `json:"mapNodes"`
}

func (h *Handler) UploadMapNodes(c *gin.Context) {
	var items NodesRequest

	err := c.Bind(&items)
	if err != nil {
		fmt.Println(err)
	}

	// err := c.Bind(&items)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// fmt.Println(items)
	// _, err := h.helper.HTTP.GetForm(c, &items)
	// if h.helper.HTTP.HandleError(c, err) {
	// 	return
	// }
	// jsonData, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// Handle error
	// }
	// fmt.Println(jsonData)
	// file, _ := json.MarshalIndent(items, "", " ")

	// _ = ioutil.WriteFile("test.json", jsonData, 0644)

	err = h.service.UploadMapNodes(items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}
