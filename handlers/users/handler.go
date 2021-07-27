package users

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/helpers"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	GetByEmail(c *gin.Context) error
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

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.repository.getAll(c)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := h.repository.get(c, c.Param("id"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, item)
}

func (h *Handler) GetByEmail(c *gin.Context) {
	item, err := h.repository.getByEmail(c, c.Param("email"))
	if err != nil || &item.Email == nil {
		c.JSON(200, nil)
	}
	if &item != nil && len(item.Email) > 0 {
		c.JSON(200, item.Email)
	}
}
