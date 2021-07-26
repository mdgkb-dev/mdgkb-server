package news

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"
	"time"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	GetBySLug(c *gin.Context) error
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	CreateLike(c *gin.Context) error
	CreateComment(c *gin.Context) error
	Delete(c *gin.Context) error
	DeleteLike(c *gin.Context) error
	DeleteComment(c *gin.Context) error
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
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &item)
	if err != nil {
		c.JSON(500, err)
	}

	err = h.uploader.Upload(c, form.File["files"][0], item.PreviewThumbnailFile.OriginalName)
	if err != nil {
		c.JSON(500, err)
	}
	item.PreviewThumbnailFile.FilenameDisk = item.PreviewThumbnailFile.OriginalName
	err = h.repository.create(c, &item)

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

func (h *Handler) CreateComment(c *gin.Context) {
	var item models.NewsComment
	err := c.ShouldBind(&item)
	if err != nil {
		c.JSON(500, err)
	}

	err = h.repository.createComment(c, &item)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, item)
}

type newsParams struct {
	PublishedOn *time.Time `form:"publishedOn"`
	Limit       int        `form:"limit"`
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

func (h *Handler) Update(c *gin.Context) {
	var item models.News
	err := c.Bind(&item)
	if err != nil {
		c.JSON(500, err)
	}
	err = h.repository.update(c, &item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
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

func (h *Handler) DeleteComment(c *gin.Context) {
	err := h.repository.deleteComment(c, c.Param("id"))
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
