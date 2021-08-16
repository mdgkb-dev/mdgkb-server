package buildings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"
)

type Handler interface {
	GetAll(c *gin.Context) error
	GetByFloorId(c *gin.Context) error
	GetById(c *gin.Context) error
	Create(c *gin.Context) error
	Delete(c *gin.Context) error
	Update(c *gin.Context) error
}

type AHandler struct {
	repository Repository
	uploader   helpers.Uploader
}

// NewHandler constructor
func NewHandler(repository Repository, uploader helpers.Uploader) *AHandler {
	return &AHandler{
		uploader:   uploader,
		repository: repository,
	}
}

func (h *AHandler) Create(c *gin.Context) {
	var item models.Building
	err := c.ShouldBindWith(&item, binding.FormMultipart)
	if err != nil {
		c.JSON(500, err)
	}

	err = h.repository.create(c, &item)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, gin.H{})
}

func (h *AHandler) GetAll(c *gin.Context) {
	buildings, err := h.repository.getAll(c)
	if err != nil {
		c.JSON(500, err)
	}
	fmt.Println(buildings)
	c.JSON(200, buildings)
}

func (h *AHandler) GetByFloorId(c *gin.Context) {
	item, err := h.repository.getByFloorId(c, c.Param("id"))
	fmt.Println(err)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, item)
}

func (h *AHandler) GetById(c *gin.Context) {
	item, err := h.repository.getById(c, c.Param("id"))
	fmt.Println(err)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, item)
}

func (h *AHandler) Delete(c *gin.Context) {
	err := h.repository.delete(c, c.Param("id"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}

func (h *AHandler) Update(c *gin.Context) {
	var building models.Building
	err := c.Bind(&building)
	if err != nil {
		c.JSON(500, err)
	}
	err = h.repository.update(c, &building)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}
