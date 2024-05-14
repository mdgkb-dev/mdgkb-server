package residencyapplications

import (
	"mdgkb/mdgkb-server/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	items, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) FTSP(c *gin.Context) {
	data, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.FTSPAnswer{Data: data, FTSP: *h.helper.SQL.ExtractFTSP(c.Request.Context())})
}
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := S.Get(c.Request.Context(), &id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) EmailExists(c *gin.Context) {
	item, err := S.EmailExists(c.Request.Context(), c.Param("email"), c.Param("courseId"))
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
	item, err := S.TypeExists(c.Request.Context(), c.Param("email"), main)
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
	err = F.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Create(c.Request.Context(), &item)
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
	err = F.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Update(c.Request.Context(), &item)
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
	err = S.UpsertMany(c.Request.Context(), items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, err)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := S.Delete(c.Request.Context(), &id)
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
	doc, err := F.FillApplicationTemplate(&item)
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
	err = F.UploadFormFiles(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.UpdateWithForm(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}
