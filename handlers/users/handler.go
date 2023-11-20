package users

import (
	"mdgkb/mdgkb-server/models"
	"net/http"

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
	item, err := h.service.Get(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetByEmail(c *gin.Context) {
	item, err := h.service.EmailExists(c.Param("email"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.User
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

func (h *Handler) Create(c *gin.Context) {
	var item models.User
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

type FavouriteForm struct {
	ID string `json:"id"`
}

func (h *Handler) AddToUser(c *gin.Context) {
	// userID, err := h.helper.Token.GetUserID(c)
	// if h.helper.HTTP.HandleError(c, err) {
	// 	return
	// }

	// domain := c.Param("domain")
	// table := fmt.Sprintf("%ss_users", domain)
	// domainCol := fmt.Sprintf("%s_id", domain)

	// fav := FavouriteForm{}
	// err = c.Bind(&fav)
	// if h.helper.HTTP.HandleError(c, err) {
	// 	return
	// }
	// domainID := fav.ID

	// values := map[string]interface{}{
	// 	domainCol: domainID,
	// 	"user_id": userID,
	// }
	// item := h.service.AddToUser(values, table)
	// if h.helper.HTTP.HandleError(c, err) {
	// 	return
	// }
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) RemoveFromUser(c *gin.Context) {
	// userID, err := h.helper.Token.GetUserID(c)
	// if h.helper.HTTP.HandleError(c, err) {
	// return
	// }

	// domain := c.Param("domain")
	// table := fmt.Sprintf("%ss_users", domain)
	// domainCol := fmt.Sprintf("%s_id", domain)

	// domainID := c.Param("id")
	// values := map[string]interface{}{
	// domainCol: domainID,
	// "user_id": userID,
	// }
	// item, err := h.service.RemoveFromUser(values, table)
	// if h.helper.HTTP.HandleError(c, err) {
	// 	return
	// }
	c.JSON(http.StatusOK, nil)
}
