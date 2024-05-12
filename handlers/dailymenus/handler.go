package dailymenus

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.DailyMenu
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = F.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Create(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) FTSP(c *gin.Context) {
	data, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.FTSPAnswer{Data: data, FTSP: *h.helper.SQL.ExtractFTSP(c.Request.Context())})
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

var u = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Ping struct {
	Ping string `json:"ping"`
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
			todayMenu, err := S.GetTodayActive(c.Request.Context())
			if err != nil {
				fmt.Printf("error sending message: %s\n", err.Error())
			}
			err = ws.WriteJSON(todayMenu)
			if err != nil {
				fmt.Printf("error sending message: %s\n", err.Error())
			}
		}
	}
}

func (h *Handler) Get(c *gin.Context) {
	item, err := S.Get(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	err := S.Delete(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.DailyMenu
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = F.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Update(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) UpdateAll(c *gin.Context) {
	var items models.DailyMenus
	err := c.Bind(&items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.UpdateAll(c.Request.Context(), items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) PDF(c *gin.Context) {
	var item models.DailyMenu
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	pdf, err := h.helper.PDF.GeneratePDF("dailyMenu", item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.Data(http.StatusOK, "application/pdf", pdf)
}

func (h *Handler) GetTodayMenu(c *gin.Context) {
	item, err := S.GetTodayActive(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}
