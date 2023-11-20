package timetables

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllWeekdays(c *gin.Context) {
	items, err := h.service.GetAllWeekdays()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}
