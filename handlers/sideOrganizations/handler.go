package sideOrganizations

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Create(c *gin.Context) error
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	Update(c *gin.Context) error
	UpdateStatus(c *gin.Context) error
	Delete(c *gin.Context) error
}

type Handler struct {
	repository IRepository
	//uploader   helper.
}

// NewHandler constructor
func NewHandler(repository IRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var item models.SideOrganization
	err := c.Bind(&item)
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = h.repository.create(c, &item)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, gin.H{})
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

func (h *Handler) Update(c *gin.Context) {
	var item models.SideOrganization
	err := c.Bind(&item)
	if err != nil {
		c.JSON(500, err)
	}
	err = h.repository.update(c, &item)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	var item models.SideOrganization
	err := c.Bind(&item)
	if err != nil {
		c.JSON(500, err)
	}
	err = h.repository.updateStatus(c, &item)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) Delete(c *gin.Context) {
	err := h.repository.delete(c, c.Param("id"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}
