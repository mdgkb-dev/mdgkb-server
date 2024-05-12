package news

import (
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FTSP(c *gin.Context) {
	data, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.FTSPAnswer{Data: data, FTSP: *h.helper.SQL.ExtractFTSP(c.Request.Context())})
}

func (h *Handler) Create(c *gin.Context) {
	var item models.News
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
	c.JSON(http.StatusOK, item)
}

func (h *Handler) RemoveTag(c *gin.Context) {
	var item models.NewsToTag
	err := c.ShouldBind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.RemoveTag(c.Request.Context(), &item)
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
	err = S.AddTag(c.Request.Context(), &item)
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

	err = S.CreateLike(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	news, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	for i := range news.News {
		news.News[i].ViewsCount = len(news.News[i].NewsViews)
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) GetMain(c *gin.Context) {
	news, err := S.GetMain(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	for i := range news.News {
		news.News[i].ViewsCount = len(news.News[i].NewsViews)
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) GetSubMain(c *gin.Context) {
	news, err := S.GetSubMain(c.Request.Context())
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

func (h *Handler) Delete(c *gin.Context) {
	err := S.Delete(c.Request.Context(), c.Param("id"))
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

	err = S.CreateComment(c.Request.Context(), &item)
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
	err = S.UpdateComment(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) RemoveComment(c *gin.Context) {
	err := S.RemoveComment(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) DeleteLike(c *gin.Context) {
	err := S.DeleteLike(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetBySLug(c *gin.Context) {
	item, err := S.GetBySlug(c.Request.Context(), c.Param("slug"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetNewsComments(c *gin.Context) {
	item, err := S.GetNewsComments(c, c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetSuggestionNews(c *gin.Context) {
	items, err := S.GetSuggestionNews(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}
