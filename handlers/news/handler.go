package news

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"
	"time"
)

type Handler interface {
	GetAll(c *gin.Context) error
	Create(c *gin.Context) error
	Delete(c *gin.Context) error
	UpdateStatus(c *gin.Context) error
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
	var item models.News
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

type newsParams struct {
	PublishedOn *time.Time `form:"publishedOn"`
}

func (h *AHandler) GetAll(c *gin.Context) {
	var newsParams newsParams
	err := c.BindQuery(&newsParams)
	if err != nil {
		c.JSON(500, err)
	}

	news, err := h.repository.getAll(c, &newsParams)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, news)
}

func (h *AHandler) UpdateStatus(c *gin.Context) {
	var item models.News
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

func (h *AHandler) Delete(c *gin.Context) {
	err := h.repository.delete(c, c.Param("id"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}
