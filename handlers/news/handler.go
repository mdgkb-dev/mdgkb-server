package news

import (
	"encoding/json"
	"fmt"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"mdgkb/mdgkb-server/models"
	"time"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	GetBySLug(c *gin.Context) error
	GetByMonth(c *gin.Context) error
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	CreateLike(c *gin.Context) error
	AddTag(c *gin.Context) error
	RemoveTag(c *gin.Context) error
	Delete(c *gin.Context) error
	DeleteLike(c *gin.Context) error
	CreateComment(c *gin.Context) error
	DeleteComment(c *gin.Context) error
	UpdateComment(c *gin.Context) error
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
		fmt.Println(err)
		c.JSON(500, err)
	}

	err = h.uploader.Upload(c, form.File["mainImage"][0], item.MainImage.FileSystemPath)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	err = h.uploader.Upload(c, form.File["previewFile"][0], item.FileInfo.FileSystemPath)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	for i, file := range form.File["gallery"] {
		err = h.uploader.Upload(c, file, item.NewsImagesNames[i])
		if err != nil {
			fmt.Println(err)
			c.JSON(500, err)
		}
	}

	err = h.repository.create(c, &item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, item)
}

func (h *Handler) RemoveTag(c *gin.Context) {
	var item models.NewsToTag
	err := c.ShouldBind(&item)
	if err != nil {
		c.JSON(500, err)
	}
	err = h.repository.removeTag(c, &item)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, item)
}

func (h *Handler) AddTag(c *gin.Context) {
	var item models.NewsToTag
	err := c.ShouldBind(&item)
	if err != nil {
		c.JSON(500, err)
	}
	err = h.repository.addTag(c, &item)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, item)
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
	Limit       int        `form:"limit"`
	FilterTags  string     `form:"filterTags"`
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
	for i := range news {
		news[i].ViewsCount = len(news[i].NewsViews)
	}
	c.JSON(200, news)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.News
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &item)
	if err != nil {
		c.JSON(500, err)
	}

	if len(form.File["mainImage"]) > 0 {
		err = h.uploader.Upload(c, form.File["mainImage"][0], item.MainImage.FileSystemPath)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, err)
		}
	}

	if len(form.File["previewFile"]) > 0 {
		err = h.uploader.Upload(c, form.File["previewFile"][0], item.FileInfo.FileSystemPath)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, err)
		}
	}

	for i, file := range form.File["gallery"] {
		err = h.uploader.Upload(c, file, item.NewsImagesNames[i])
		if err != nil {
			fmt.Println(err)
			c.JSON(500, err)
		}
	}

	err = h.repository.update(c, &item)
	if err != nil {
		fmt.Println(err)
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

func (h *Handler) RemoveComment(c *gin.Context) {
	err := h.repository.removeComment(c, c.Param("id"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
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

func (h *Handler) UpdateComment(c *gin.Context) {
	var item models.NewsComment
	err := c.Bind(&item)
	err = h.repository.updateComment(c, &item)
	if err != nil {
		fmt.Println(err)
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
	item.ViewsCount = len(item.NewsViews)
	ip, err := httpHelper.GetClientIPHelper(c.Request)
	newsView := models.NewsViews{IPAddress: ip, NewsID: item.ID}
	err = h.repository.createViewOfNews(c, &newsView)

	c.JSON(200, item)
}

type monthParams struct {
	Month int `form:"month"`
	Year  int `form:"year"`
}

func (h *Handler) GetByMonth(c *gin.Context) {
	var monthParams monthParams
	err := c.BindQuery(&monthParams)
	if err != nil {
		c.JSON(500, err)
	}

	news, err := h.repository.getByMonth(c, &monthParams)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, news)
}
