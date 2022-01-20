package users

import (
	"fmt"
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll()
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

func (h *Handler) GetByEmail(c *gin.Context) {
	item, err := h.service.EmailExists(c.Param("email"))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.User
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)

	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

type FavouriteForm struct {
	ID string `json:"id"`
}

func (h *Handler) AddToUser(c *gin.Context) {
	userID, err := h.helper.Token.GetUserID(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
		return
	}

	domain := c.Param("domain")
	table := fmt.Sprintf("%ss_users", domain)
	domainCol := fmt.Sprintf("%s_id", domain)

	fav := FavouriteForm{}
	err = c.Bind(&fav)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	domainID := fav.ID

	values := map[string]interface{}{
		domainCol: domainID,
		"user_id": userID,
	}
	item := h.service.AddToUser(values, table)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) RemoveFromUser(c *gin.Context) {
	userID, err := h.helper.Token.GetUserID(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
		return
	}

	domain := c.Param("domain")
	table := fmt.Sprintf("%ss_users", domain)
	domainCol := fmt.Sprintf("%s_id", domain)

	domainID := c.Param("id")
	values := map[string]interface{}{
		domainCol: domainID,
		"user_id": userID,
	}
	item := h.service.RemoveFromUser(values, table)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}
