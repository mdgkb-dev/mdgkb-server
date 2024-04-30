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
	err = F.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Create(c.Request.Context(), &item)
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

func (h *Handler) FTSP(c *gin.Context) {
	data, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.FTSPAnswer{Data: data, FTSP: *h.helper.SQL.ExtractFTSP(c.Request.Context())})
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := S.Get(c.Request.Context(), id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := S.Delete(c.Request.Context(), id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Question
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Update(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) ReadAnswers(c *gin.Context) {
	userID := c.Param("user-id")
	err := S.ReadAnswers(c.Request.Context(), userID)
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
	err = S.ChangeNewStatus(c.Request.Context(), id, isNew)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Publish(c *gin.Context) {
	id := c.Param("id")
	err := S.Publish(c, id)
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
	err = S.UpsertMany(c, items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, err)
}
