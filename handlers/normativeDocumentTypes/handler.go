package normativeDocumentTypes

import (
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(c *gin.Context) error
}

type AHandler struct {
	repository IRepository
	uploader   helpers.Uploader
}

func NewHandler(repository IRepository, uploader helpers.Uploader) *AHandler {
	return &AHandler{
		uploader:   uploader,
		repository: repository,
	}
}

func (h *AHandler) Create(c *gin.Context) {
	var item models.NormativeDocumentType
	err := c.Bind(&item)
	if err != nil {
		c.JSON(500, err)
	}

	err = h.repository.create(c, &item)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, gin.H{})
}
