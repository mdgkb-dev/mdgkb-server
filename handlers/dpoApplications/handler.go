package dpoApplications

import (
	"github.com/gin-gonic/gin"
	"io"
	"mdgkb/mdgkb-server/models"
	"net/http"
	"time"
)

func (h *Handler) GetAll(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	items, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(&id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) EmailExists(c *gin.Context) {
	item, err := h.service.EmailExists(c.Param("email"), c.Param("courseId"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Create(c *gin.Context) {
	var item models.DpoApplication

	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Create(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	h.sse.Message <- "New application"
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.DpoApplication
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	//if item.FormValue.EmailNotify {
	//	body, err := h.helper.Templater.ParseTemplate(item, "email/application_update_status.gohtml")
	//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	//		return
	//	}
	//	err = h.helper.Email.SendEmail([]string{item.FormValue.User.Email}, "Статус вашей заявки обновлён", body)
	//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	//		return
	//	}
	//}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(&id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) SubscribeCreate(c *gin.Context) {
	go func() {
		h.sse.Message <- "ping"
		for {
			time.Sleep(time.Second * 60)
			h.sse.Message <- "ping"
		}
	}()

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-h.sse.Message; ok {
			c.SSEvent("dpoApplicationCreate", msg)
			return true
		}
		return false
	})
}
