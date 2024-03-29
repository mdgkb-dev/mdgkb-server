package doctors

import (
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//type doctorsParams struct {
//	Main  bool `form:"main"`
//	Limit int  `form:"limit"`
//}

func (h *Handler) Create(c *gin.Context) {
	var item models.Doctor
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
	c.JSON(http.StatusOK, gin.H{})
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

func (h *Handler) GetAllAdmin(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	items, err := h.service.GetAllAdmin()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := h.service.Get(c.Param("slug"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetByDivisionID(c *gin.Context) {
	item, err := h.service.GetByDivisionID(c.Param("divisionId"))
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

func (h *Handler) Update(c *gin.Context) {
	var item models.Doctor
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

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) CreateComment(c *gin.Context) {
	var item models.DoctorComment
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
	var item models.DoctorComment
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) CreateSlugs(c *gin.Context) {
	err := h.service.CreateSlugs()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) Search(c *gin.Context) {
	query := c.Query("query")
	if query != "" {
		items, err := h.service.Search(query)
		if h.helper.HTTP.HandleError(c, err) {
			return
		}
		c.JSON(http.StatusOK, items)
		return
	}
	c.JSON(http.StatusOK, nil)
}
