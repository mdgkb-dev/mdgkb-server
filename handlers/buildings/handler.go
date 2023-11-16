package buildings

import (
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Building
	err := c.ShouldBindWith(&item, binding.FormMultipart)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.repository.create(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetAll(c *gin.Context) {
	buildings, err := h.repository.getAll(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, buildings)
}

func (h *Handler) GetByFloorID(c *gin.Context) {
	item, err := h.repository.getByFloorID(c, c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetByID(c *gin.Context) {
	item, err := h.repository.getByID(c, c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	err := h.repository.delete(c, c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var building models.Building
	err := c.Bind(&building)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.repository.update(c, &building)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
