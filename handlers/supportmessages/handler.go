package supportmessages

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.SupportMessage
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Create(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	emailStruct := struct {
		SupportMessage *models.SupportMessage
		Host           string
	}{
		&item,
		h.helper.HTTP.Host,
	}
	mail, err := h.helper.Templater.ParseTemplate(emailStruct, "email/newSupportMessage.gohtml")
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.helper.Email.SendEmail([]string{"lakkinzimusic@gmail.com"}, fmt.Sprintf("Новый вопрос: %s", item.Theme), mail)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	h.helper.Broker.SendEvent("support-message-create", item)
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	items, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.SupportMessage
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ChangeNewStatus(c *gin.Context) {
	id := c.Param("id")
	isNew, err := strconv.ParseBool(c.Query("isNew"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.ChangeNewStatus(id, isNew)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
