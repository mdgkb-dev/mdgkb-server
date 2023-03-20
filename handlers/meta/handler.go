package meta

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pro-assistance/pro-assister/projecthelper"
)

func (h *Handler) GetCount(c *gin.Context) {
	table := c.Param("table")
	items, err := h.service.GetCount(&table)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetSchema(c *gin.Context) {
	c.JSON(http.StatusOK, projecthelper.SchemasLib)
}

func (h *Handler) GetSocial(c *gin.Context) {
	c.JSON(http.StatusOK, h.helper.Social.GetWebFeed())
}

func (h *Handler) GetApplicationsCounts(c *gin.Context) {
	items, err := h.service.GetApplicationsCounts()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetOptions(c *gin.Context) {
	var optionModel models.OptionModel
	err := c.BindQuery(&optionModel)
	optionModel.TableName = c.Query("tableName")
	optionModel.Label = c.Query("label")
	optionModel.Value = c.Query("value")
	optionModel.SortColumn = c.Query("sortColumn")
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	items, err := h.service.GetOptions(&optionModel)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

var u = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) GetWeb(c *gin.Context) {
	u.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := u.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		_, m, err := ws.ReadMessage()
		if err == nil && string(m) == "ping" {
			items, err := h.service.GetApplicationsCounts()
			if err != nil {
				fmt.Printf("error sending message: %s\n", err.Error())
			}
			err = ws.WriteJSON(items)
			if err != nil {
				fmt.Printf("error sending message: %s\n", err.Error())
			}
		}
	}
}
