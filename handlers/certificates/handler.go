package certificates

import (
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

type WithDeleted struct {
	Certificates          models.Certificates `json:"certificates"`
	CertificatesForDelete []uuid.UUID         `json:"certificatesForDelete"`
}

func (h *Handler) UpdateMany(c *gin.Context) {
	var item WithDeleted
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, item.Certificates, files)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.UpsertMany(item.Certificates)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.DeleteMany(item.CertificatesForDelete)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item.Certificates)
}
