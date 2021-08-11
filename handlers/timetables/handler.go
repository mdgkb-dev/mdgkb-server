package timetables

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	"mdgkb/mdgkb-server/helpers"
)

type IHandler interface {
	GetAllWeekdays(c *gin.Context) error
}

type Handler struct {
	repository IRepository
	uploader   helpers.Uploader
}

// NewHandler constructor
func NewHandler(repository IRepository, uploader helpers.Uploader) *Handler {
	return &Handler{
		uploader:   uploader,
		repository: repository,
	}
}

func (h *Handler) GetAllWeekdays(c *gin.Context) {
	items, err := h.repository.getAllWeekdays(c)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, items)
}
