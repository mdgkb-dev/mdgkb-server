package dataexport

import (
	"encoding/json"
	"fmt"
	"mdgkb/mdgkb-server/handlers/news"
	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/models/exportmodels"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Export(c *gin.Context) {
}

func (h *Handler) Data(c *gin.Context) {
	opts, err := models.ExportOptions{}.New(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = json.Unmarshal([]byte(c.Query("exportOptions")), &opts)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fmt.Println(opts)
	newsExport := exportmodels.NewsView{}

	err = opts.Parse(&newsExport)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fmt.Println(opts)
	items, err := news.S.GetAggregateViews(c.Request.Context(), &newsExport)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	item := models.ChartQuery{}
	item.InitFromDataSets(items)
	c.JSON(http.StatusOK, item)
}
