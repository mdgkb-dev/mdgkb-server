package preparations

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"net/http"

	"github.com/gin-gonic/gin"

	"mdgkb/mdgkb-server/models"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Preparation
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)

	err = h.service.Create(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := h.service.Get(httpHelper.GetID(c))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	err := h.service.Delete(httpHelper.GetID(c))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Preparation
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)

	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

type PreparationsWithDeleted struct {
	Preparations           models.Preparations `json:"preparations"`
	PreparationsForDeleted []uuid.UUID         `json:"preparationsForDeleted"`
}

func (h *Handler) UpdateMany(c *gin.Context) {
	var items PreparationsWithDeleted
	_, err := h.helper.HTTP.GetForm(c, &items)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	//err = h.filesService.Upload(c, &item, files)

	err = h.service.UpsertMany(items)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items.Preparations)
}

func (h *Handler) GetTags(c *gin.Context) {
	items, err := h.service.GetTags()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}
