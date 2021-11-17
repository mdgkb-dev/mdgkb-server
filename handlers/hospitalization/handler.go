package hospitalization

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) PDF(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	fmt.Println("item =====>", item)
	pdf, err := h.helper.PDF.GeneratePDF("hospitalizationDocList", item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=response")
	c.Data(http.StatusOK, "application/pdf", pdf)
}
