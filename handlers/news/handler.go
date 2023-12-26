package news

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FTSP(c *gin.Context) {
	var item models.FTSPQuery
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fmt.Println(item)
	err = h.service.SetQueryFilter(c, item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	data, err := h.service.GetAll(true)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.FTSPAnswer{Data: data})
}

func (h *Handler) Create(c *gin.Context) {
	var item models.News
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
	c.JSON(http.StatusOK, item)
}

func (h *Handler) RemoveTag(c *gin.Context) {
	var item models.NewsToTag
	err := c.ShouldBind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.RemoveTag(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) AddTag(c *gin.Context) {
	var item models.NewsToTag
	err := c.ShouldBind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.AddTag(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) CreateLike(c *gin.Context) {
	var item models.NewsLike
	err := c.ShouldBind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.service.CreateLike(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	err := h.service.SetQueryFilter(c, models.FTSPQuery{})
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	news, err := h.service.GetAll(false)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	for i := range news.News {
		news.News[i].ViewsCount = len(news.News[i].NewsViews)
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) GetMain(c *gin.Context) {
	err := h.service.SetQueryFilter(c, models.FTSPQuery{})
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	news, err := h.service.GetMain()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	for i := range news.News {
		news.News[i].ViewsCount = len(news.News[i].NewsViews)
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) GetSubMain(c *gin.Context) {
	err := h.service.SetQueryFilter(c, models.FTSPQuery{})
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	news, err := h.service.GetSubMain()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	for i := range news.News {
		news.News[i].ViewsCount = len(news.News[i].NewsViews)
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.News
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

func (h *Handler) Delete(c *gin.Context) {
	err := h.service.Delete(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) CreateComment(c *gin.Context) {
	var item models.NewsComment
	err := c.ShouldBind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.service.CreateComment(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) UpdateComment(c *gin.Context) {
	var item models.NewsComment
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.UpdateComment(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) RemoveComment(c *gin.Context) {
	err := h.service.RemoveComment(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) DeleteLike(c *gin.Context) {
	err := h.service.DeleteLike(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetBySLug(c *gin.Context) {
	item, err := h.service.GetBySlug(c, c.Param("slug"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetNewsComments(c *gin.Context) {
	item, err := h.service.GetNewsComments(c, c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetSuggestionNews(c *gin.Context) {
	items, err := h.service.GetSuggestionNews(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}
