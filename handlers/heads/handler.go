package heads

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Head
	_, err := httpHelper.GetForm(c, &item)
	//if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//err = h.filesService.Upload(c, &item, files)

	err = h.service.Create(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll()
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := h.service.Get(c.Param("id"))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}



func (h *Handler) Delete(c *gin.Context) {
	err := h.service.Delete(c.Param("id"))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Head
	files, err := httpHelper.GetForm(c, &item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)

	err = h.service.Update(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
