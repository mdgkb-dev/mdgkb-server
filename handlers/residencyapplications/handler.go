package residencyapplications

import (
	"mdgkb/mdgkb-server/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
	item, err := h.service.Get(&id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) EmailExists(c *gin.Context) {
	item, err := h.service.EmailExists(c.Param("email"), c.Param("courseId"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) TypeExists(c *gin.Context) {
	main, err := strconv.ParseBool(c.Param("main"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	item, err := h.service.TypeExists(c.Param("email"), main)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Create(c *gin.Context) {
	var item models.ResidencyApplication

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
	h.helper.Broker.SendEvent("residency-application-create", item)
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.ResidencyApplication
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) UpsertMany(c *gin.Context) {
	var items models.ResidencyApplications
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

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(&id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) FillApplicationTemplate(c *gin.Context) {
	var item models.ResidencyApplication
	err := c.Bind(&item)
	t1 := item.FormValue.User.Human.DateBirth.Add(time.Hour * 3)
	item.FormValue.User.Human.DateBirth = &t1
	item.FormValue.NormalizeDateFields()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	doc, err := h.filesService.FillApplicationTemplate(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.wordprocessingml.document", doc)
}

func (h *Handler) UpdateWithForm(c *gin.Context) {
	var item models.FormValue
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.UploadFormFiles(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.UpdateWithForm(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}
