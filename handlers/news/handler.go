package news

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"mdgkb/mdgkb-server/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.News
	files, err := httpHelper.GetForm(c, &item)
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
	c.JSON(http.StatusOK, item)
}

func (h *Handler) RemoveTag(c *gin.Context) {
	var item models.NewsToTag
	err := c.ShouldBind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.RemoveTag(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) AddTag(c *gin.Context) {
	var item models.NewsToTag
	err := c.ShouldBind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.AddTag(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) CreateLike(c *gin.Context) {
	var item models.NewsLike
	err := c.ShouldBind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	err = h.service.CreateLike(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, item)
}

type newsParams struct {
	PublishedOn *time.Time `form:"publishedOn"`
	Limit       int        `form:"limit"`
	FilterTags  string     `form:"filterTags"`
	OrderByView string     `form:"orderByView"`
	Events      bool       `form:"events"`
}

func (h *Handler) GetAll(c *gin.Context) {
	var newsParams newsParams
	err := c.BindQuery(&newsParams)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	news, err := h.service.GetAll(&newsParams)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	for i := range news {
		news[i].ViewsCount = len(news[i].NewsViews)
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) GetAllAdmin(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	news, err := h.service.GetAllAdmin()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) GetAllRelationsNews(c *gin.Context) {
	var newsParams newsParams
	err := c.BindQuery(&newsParams)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	news, err := h.service.GetAllRelationsNews(&newsParams)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	for i := range news {
		news[i].ViewsCount = len(news[i].NewsViews)
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.News
	files, err := httpHelper.GetForm(c, &item)
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
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	err := h.service.Delete(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) CreateComment(c *gin.Context) {
	var item models.NewsComment
	err := c.ShouldBind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	err = h.service.CreateComment(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) UpdateComment(c *gin.Context) {
	var item models.NewsComment
	err := c.Bind(&item)
	err = h.service.UpdateComment(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(200, item)
}

func (h *Handler) RemoveComment(c *gin.Context) {
	err := h.service.RemoveComment(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) DeleteLike(c *gin.Context) {
	err := h.service.DeleteLike(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetBySLug(c *gin.Context) {
	item, err := h.service.GetBySlug(c.Param("slug"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	item.ViewsCount = len(item.NewsViews)
	ip, err := httpHelper.GetClientIPHelper(c.Request)
	newsView := models.NewsView{IPAddress: ip, NewsID: item.ID}
	err = h.service.CreateViewOfNews(&newsView)

	c.JSON(http.StatusOK, item)
}

type monthParams struct {
	Month int `form:"month"`
	Year  int `form:"year"`
}

func (h *Handler) GetByMonth(c *gin.Context) {
	var monthParams monthParams
	err := c.BindQuery(&monthParams)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	items, err := h.service.GetByMonth(&monthParams)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, items)
}
