package paidPrograms

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"net/http"

	"github.com/gin-gonic/gin"

	"mdgkb/mdgkb-server/models"
)


func (h *Handler) Get(c *gin.Context) {
	item, err := h.service.Get(httpHelper.GetID(c))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.PaidProgram
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	//err = h.filesService.(c, &item, files)

	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
