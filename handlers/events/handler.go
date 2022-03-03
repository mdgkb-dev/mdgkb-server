package events

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
	"net/http"
)

func (h *Handler) CreateEventApplication(c *gin.Context) {
	var item models.EventApplication
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	userID, err := h.helper.Token.GetUserID(c)
	err = h.service.CreateEventApplication(&item)
	item.UserID = *userID
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) EventApplicationsPDF(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	pdf, err := h.helper.PDF.GeneratePDF("eventApplications", item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	h.helper.HTTP.SetFileHeaders(c, "Заявки на мероприятие")
	c.Data(http.StatusOK, "application/pdf", pdf)
}

func (h *Handler) GetAllForMain(c *gin.Context) {
	items, err := h.service.GetAllForMain()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}
