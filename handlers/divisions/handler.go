package divisions

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	Create(c *gin.Context) error
	Delete(c *gin.Context) error
	Update(c *gin.Context) error
	CreateComment(c *gin.Context) error
	UpdateComment(c *gin.Context) error
	RemoveComment(c *gin.Context) error
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
	var item models.Division
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
		return
	}

	for i, file := range form.File["gallery"] {
		err = h.uploader.Upload(c, file, item.DivisionImagesNames[i])
		if err != nil {
			fmt.Println(err)
			c.JSON(500, err)
		}
	}

	err = h.repository.create(c, &item)
	if err != nil {
		fmt.Println(err)
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

func (h *Handler) Delete(c *gin.Context) {
	err := h.repository.delete(c, c.Param("id"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Division
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &item)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	for i, file := range form.File["gallery"] {
		err = h.uploader.Upload(c, file, item.DivisionImagesNames[i])
		if err != nil {
			fmt.Println(err)
			c.JSON(500, err)
		}
	}

	err = h.repository.update(c, &item)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) CreateComment(c *gin.Context) {
	var item models.DivisionComment
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
	var item models.DivisionComment
	err := c.Bind(&item)
	err = h.repository.updateComment(c, &item)
	if err != nil {
		fmt.Println(err)
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
