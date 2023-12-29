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

func (h *Handler) Create(c *gin.Context) {
	var items NodesRequest

	_, err := h.helper.HTTP.GetForm(c, &items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err := c.Bind(&items)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	//	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// Handle error
	// }
	// fmt.Println(jsonData)
	// file, _ := json.MarshalIndent(items, "", " ")

	//	_ = ioutil.WriteFile("test.json", jsonData, 0644)

	err = h.service.UploadMapNodes(items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, nil)
}
