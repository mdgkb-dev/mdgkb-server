package comments

import (
	"mdgkb/mdgkb-server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	comments, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *Handler) GetAllMain(c *gin.Context) {
	comments, err := h.service.GetAllMain()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *Handler) UpdateOne(c *gin.Context) {
	var item models.Comment
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.UpdateOne(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) UpsertOne(c *gin.Context) {
	var item models.Comment
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.UpsertOne(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}
