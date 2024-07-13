package users

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/pro-assistance/pro-assister/middleware"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	items, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := S.Get(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) FTSP(c *gin.Context) {
	data, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.FTSPAnswer{Data: data, FTSP: *h.helper.SQL.ExtractFTSP(c.Request.Context())})
}

func (h *Handler) GetByEmail(c *gin.Context) {
	item, err := S.EmailExists(c.Request.Context(), c.Param("email"))
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

func (h *Handler) Create(c *gin.Context) {
	var item models.User
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

type FavouriteForm struct {
	ID string `json:"id"`
}

func (h *Handler) AddToUser(c *gin.Context) {
	userID, err := h.helper.Token.ExtractTokenMetadata(c.Request, middleware.ClaimUserID)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	domain := c.Param("domain")
	table := fmt.Sprintf("%ss_users", domain)
	domainCol := fmt.Sprintf("%s_id", domain)

	fav := FavouriteForm{}
	err = c.Bind(&fav)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	domainID := fav.ID

	values := map[string]interface{}{
		domainCol: domainID,
		"user_id": userID,
	}
	item := S.AddToUser(c.Request.Context(), values, table)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) RemoveFromUser(c *gin.Context) {
	userID, err := h.helper.Token.ExtractTokenMetadata(c.Request, middleware.ClaimUserID)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	domain := c.Param("domain")
	table := fmt.Sprintf("%ss_users", domain)
	// domainCol := fmt.Sprintf("%s_id", domain)

	_ = c.Param("id")
	values := map[string]interface{}{
		// domainCol: domainID,
		"user_id": userID,
	}
	item := S.RemoveFromUser(c.Request.Context(), values, table)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}
