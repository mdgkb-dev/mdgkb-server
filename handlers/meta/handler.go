package meta

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type Address struct {
	ID       string      `json:"id"`
	FullName string      `json:"fullName"`
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Zip      interface{} `json:"zip"`
}

type Addresses []*Address

type Result struct {
	Result Addresses `json:"result"`
}

func (h *Handler) GetAddress(c *gin.Context) {
	var item models.KladrAPI
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	ctx := context.Background()
	url := item.GetURL()
	fmt.Println(url)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}
	cli := &http.Client{}
	resp, err := cli.Do(request)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()
	res := Result{}
	err = json.Unmarshal(body, &res)
	fmt.Println(err)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, res.Result[1:])
}
