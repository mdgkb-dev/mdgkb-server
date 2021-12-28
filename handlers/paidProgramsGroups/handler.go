package paidProgramsGroups

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"net/http"

	"github.com/gin-gonic/gin"

	"mdgkb/mdgkb-server/models"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.PaidProgramsGroup
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
	var item models.PaidProgramsGroup
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

func (h *Handler) UpdateMany(c *gin.Context) {
	var items models.PaidProgramsGroupsStruct
	_, err := h.helper.HTTP.GetForm(c, &items)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	//err = h.filesService.Upload(c, &item, files)

	err = h.service.UpsertMany(items)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items.PaidProgramsGroups)
}
