package questions

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Question
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
		Question *models.Question
		Host     string
	}{
		&item,
		h.helper.HTTP.Host,
	}
	mail, err := h.helper.Templater.ParseTemplate(emailStruct, "email/newQuestion.gohtml")
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.helper.Email.SendEmail([]string{"lakkinzimusic@gmail.com"}, fmt.Sprintf("Новый вопрос: %s", item.Theme), mail)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	h.helper.Broker.SendEvent("question-create", item)
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
	var item models.Question
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

func (h *Handler) ReadAnswers(c *gin.Context) {
	userID := c.Param("user-id")
	err := h.service.ReadAnswers(userID)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
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

func (h *Handler) Publish(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Publish(id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) UpsertMany(c *gin.Context) {
	var items models.Questions
	err := c.Bind(&items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.UpsertMany(items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, err)
}
