package postgraduateDocumentTypes

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Create(c *gin.Context) {
	var item models.PostgraduateDocumentTypes
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	//err = h.service.UpsertMany(item)
	//if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	c.JSON(http.StatusOK, item)
}

type PostgraduateDocumentTypesWithDelete struct {
	PostgraduateDocumentTypes          models.PostgraduateDocumentTypes `json:"postgraduateDocumentTypes"`
	PostgraduateDocumentTypesForDelete []uuid.UUID                      `json:"postgraduateDocumentTypesForDelete"`
}

func (h *Handler) Update(c *gin.Context) {
	var item PostgraduateDocumentTypesWithDelete
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item.PostgraduateDocumentTypes, files)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.UpsertMany(item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
