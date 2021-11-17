package users

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) GetByEmail(c *gin.Context) {
	item, err := h.service.EmailExists(c.Param("email"))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}
