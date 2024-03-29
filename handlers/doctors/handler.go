package doctors

import (
	"net/http"

	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Doctor
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
	c.JSON(http.StatusOK, gin.H{})
}

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
	item, err := S.Get(c.Request.Context(), c.Param("slug"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetByDivisionID(c *gin.Context) {
	item, err := S.GetByDivisionID(c.Request.Context(), c.Param("divisionId"))
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

func (h *Handler) Update(c *gin.Context) {
	var item models.Doctor
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

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) CreateComment(c *gin.Context) {
	var item models.DoctorComment
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
	var item models.DoctorComment
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) CreateSlugs(c *gin.Context) {
	err := S.CreateSlugs(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, nil)
}
