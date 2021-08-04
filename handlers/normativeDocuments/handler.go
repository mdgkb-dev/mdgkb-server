package normativeDocuments

import (
	"encoding/json"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kennygrant/sanitize"
)

type IHandler interface {
	Create(c *gin.Context) error
	Get(c *gin.Context) error
	GetAll(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type Handler struct {
	repository IRepository
	uploader   helpers.Uploader
}

func NewHandler(repository IRepository, uploader helpers.Uploader) *Handler {
	return &Handler{
		uploader:   uploader,
		repository: repository,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var document models.NormativeDocument
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &document)

	if err != nil {
		c.JSON(500, err)
		return
	}

	document.NormativeDocumentTypeId = document.NormativeDocumentType.ID

	if document.FileInfo != nil {
		document.FileInfo.ID = uuid.New()
		document.FileInfo.FileSystemPath = "normative-documents" + "/" + document.FileInfo.ID.String() + ".pdf"
		document.FileInfo.OriginalName = sanitize.BaseName(document.FileInfo.OriginalName)
		err = h.uploader.Upload(c, form.File["files"][0], document.FileInfo.FileSystemPath)

		if err != nil {
			c.JSON(500, err)
			return
		}
	}

	err = h.repository.create(c, &document)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, gin.H{})
}

func (h *Handler) Get(c *gin.Context) {
	item, err := h.repository.get(c, c.Param("id"))

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.repository.getAll(c)

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, items)
}

func (h *Handler) Update(c *gin.Context) {
	var document models.NormativeDocument
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &document)

	if err != nil {
		c.JSON(500, err)
		return
	}

	document.NormativeDocumentTypeId = document.NormativeDocumentType.ID

	if len(form.File["files"]) > 0 {
		// TODO: Переделать на замену файла, вместо создания нового.

		if document.FileInfo.ID == uuid.Nil {
			document.FileInfo.ID = uuid.New()
		}

		document.FileInfo.FileSystemPath = "normative-documents" + "/" + document.FileInfo.ID.String() + ".pdf"
		document.FileInfo.OriginalName = sanitize.BaseName(document.FileInfo.OriginalName)
		err = h.uploader.Upload(c, form.File["files"][0], document.FileInfo.FileSystemPath)

		if err != nil {
			c.JSON(500, err)
			return
		}
	}

	err = h.repository.update(c, &document)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, gin.H{})
}

func (h *Handler) Delete(c *gin.Context) {
	// TODO: Добавить возможность удаления файла.
	err := h.repository.delete(c, c.Param("id"))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, gin.H{})
}
