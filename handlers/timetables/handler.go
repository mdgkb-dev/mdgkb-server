package timetables

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllWeekdays(c *gin.Context) {
	items, err := h.service.GetAllWeekdays()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, items)
}
