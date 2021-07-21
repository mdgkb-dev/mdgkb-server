package news

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"
	"time"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	GetBySLug(c *gin.Context) error
	Create(c *gin.Context) error
	CreateLike(c *gin.Context) error
	Delete(c *gin.Context) error
	DeleteLike(c *gin.Context) error
	UpdateStatus(c *gin.Context) error
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

func (h *Handler) Create(c *gin.Context) {
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

func (h *Handler) CreateLike(c *gin.Context) {
	var item models.NewsLike
	err := c.ShouldBind(&item)
	if err != nil {
		c.JSON(500, err)
	}

	err = h.repository.createLike(c, &item)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, item)
}

type newsParams struct {
	PublishedOn *time.Time `form:"publishedOn"`
}

func (h *Handler) GetAll(c *gin.Context) {
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

func (h *Handler) UpdateStatus(c *gin.Context) {
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

func (h *Handler) Delete(c *gin.Context) {
	err := h.repository.delete(c, c.Param("id"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) DeleteLike(c *gin.Context) {
	err := h.repository.deleteLike(c, c.Param("id"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) GetBySLug(c *gin.Context) {
	item, err := h.repository.getBySlug(c, c.Param("slug"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, item)
}
