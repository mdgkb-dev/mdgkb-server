package events

import (
	"net/http"

	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateEventApplication(c *gin.Context) {
	var item models.EventApplication
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	userID, err := h.helper.Token.ExtractTokenMetadata(c.Request, middleware.ClaimUserID)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.CreateEventApplication(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	item.UserID = uuid.MustParse(userID)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) EventApplicationsPDF(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	pdf, err := h.helper.PDF.GeneratePDF("eventApplications", item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	h.helper.HTTP.SetFileHeaders(c, "Заявки на мероприятие")
	c.Data(http.StatusOK, "application/pdf", pdf)
}

func (h *Handler) GetAllForMain(c *gin.Context) {
	items, err := h.service.GetAllForMain()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}
