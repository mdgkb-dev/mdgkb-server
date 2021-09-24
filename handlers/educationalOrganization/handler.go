package educationalOrganization

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


//
func (h *Handler) Get(c *gin.Context) {
	item, err := h.service.Get()
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.EducationalOrganization
	err := c.Bind(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Update(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}
