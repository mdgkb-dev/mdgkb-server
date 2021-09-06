package banners

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
	Create(c *gin.Context) error
	Delete(c *gin.Context) error
	Update(c *gin.Context) error
	UpdateAllOrder(c *gin.Context) error
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
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := h.repository.get(c, c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, item)
}

func (h *Handler) Create(c *gin.Context) {
	var item models.Banner
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	err = h.uploader.Upload(c, form.File["banners"][0], item.FileInfo.FileSystemPath)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	err = h.repository.create(c, &item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, gin.H{})
}

func (h *Handler) Delete(c *gin.Context) {
	err := h.repository.delete(c, c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Banner
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	if len(form.File["banners"]) > 0 {
		err = h.uploader.Upload(c, form.File["banners"][0], item.FileInfo.OriginalName)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, err)
		}
		item.FileInfo.FileSystemPath = item.FileInfo.OriginalName
	}

	err = h.repository.update(c, &item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) UpdateAllOrder(c *gin.Context) {
	var item []*models.Banner
	err := c.Bind(&item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	err = h.repository.updateAllOrder(c, item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}
