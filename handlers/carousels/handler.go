package carousels

import (
	"encoding/json"
	"fmt"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	GetByKey(c *gin.Context) error
	Create(c *gin.Context) error
	Delete(c *gin.Context) error
	Update(c *gin.Context) error
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
	var item models.Carousel
	form, _ := c.MultipartForm()
	fmt.Println([]byte(form.Value["form"][0]))
	err := json.Unmarshal([]byte(form.Value["form"][0]), &item)
	if err != nil {
		c.JSON(500, err)
	}

	if len(form.File["slide"]) > 0 {
		for _, file := range form.File["slide"] {
			err = h.uploader.Upload(c, file, item.CarouselSlides[0].FileInfo.FileSystemPath)
			if err != nil {
				c.JSON(500, err)
			}
		}
	}
	err = h.repository.create(c, &item)
	if err != nil {
		c.JSON(500, err)
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

func (h *Handler) GetByKey(c *gin.Context) {
	item, err := h.repository.getByKey(c, c.Param("key"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Carousel
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &item)
	if err != nil {
		c.JSON(500, err)
	}

	if len(form.File["slide"]) > 0 {
		for i, file := range form.File["slide"] {
			err = h.uploader.Upload(c, file, item.CarouselSlidesNames[i])
			if err != nil {
				c.JSON(500, err)
			}
		}
	}
	err = h.repository.update(c, &item)
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
