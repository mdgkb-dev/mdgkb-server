package divisions

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"net/http"

	"github.com/gin-gonic/gin"

	"mdgkb/mdgkb-server/models"
)


func (h *Handler) Create(c *gin.Context) {
	var item models.Division
	files, err := httpHelper.GetForm(c, &item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)

	err = h.service.Create(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(200, gin.H{})
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll()
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(200, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := h.service.Get(httpHelper.GetID(c))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(200, item)
}

func (h *Handler) Delete(c *gin.Context) {
	err := h.service.Delete(httpHelper.GetID(c))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Division
	files, err := httpHelper.GetForm(c, &item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)

	err = h.service.Update(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) CreateComment(c *gin.Context) {
	var item models.DivisionComment
	err := c.ShouldBind(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.CreateComment(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(200, item)
}

func (h *Handler) UpdateComment(c *gin.Context) {
	var item models.DivisionComment
	err := c.Bind(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.UpdateComment( &item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) RemoveComment(c *gin.Context) {
	err := h.service.RemoveComment(httpHelper.GetID(c))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(200, gin.H{})
}
