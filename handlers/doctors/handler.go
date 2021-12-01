package doctors

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type doctorsParams struct {
	Limit int `form:"limit"`
}

func (h *Handler) Create(c *gin.Context) {
	var item models.Doctor
	_, err := httpHelper.GetForm(c, &item)
	//if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//err = h.filesService.Upload(c, &item, files)

	err = h.service.Create(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetAll(c *gin.Context) {
	var doctorsParams doctorsParams
	err := c.BindQuery(&doctorsParams)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	items, err := h.service.GetAll(&doctorsParams)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := h.service.Get(c.Param("id"))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetByDivisionID(c *gin.Context) {
	item, err := h.service.GetByDivisionID(c.Param("divisionId"))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	err := h.service.Delete(c.Param("id"))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Doctor
	files, err := httpHelper.GetForm(c, &item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)

	err = h.service.Update(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) CreateComment(c *gin.Context) {
	var item models.DoctorComment
	err := c.ShouldBind(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	err = h.service.CreateComment(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) UpdateComment(c *gin.Context) {
	var item models.DoctorComment
	err := c.Bind(&item)
	err = h.service.UpdateComment(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
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
