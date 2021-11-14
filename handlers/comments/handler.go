package comments

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"mdgkb/mdgkb-server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type commentsParams struct {
	Limit      int  `form:"limit"`
	ModChecked *bool `form:"modChecked"`
	Positive   *bool `form:"positive"`
}

func (h *Handler) GetAll(c *gin.Context) {
	var commentsParams commentsParams
	err := c.BindQuery(&commentsParams)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	comments, err := h.service.GetAll(&commentsParams)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *Handler) UpdateOne(c *gin.Context) {
	var item models.Comment
	err := c.Bind(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.UpdateOne(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}
