package meta

import (
	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/structsreader"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCount(c *gin.Context) {
	table := c.Param("table")
	items, err := h.service.GetCount(&table)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

var schemas = structsreader.GetSchemas()

func (h *Handler) GetSchema(c *gin.Context) {
	c.JSON(http.StatusOK, schemas)
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
