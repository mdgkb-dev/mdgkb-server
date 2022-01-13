package donorRules

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	userID, err := h.helper.Token.GetUserID(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
		return
	}
	items, err := h.service.GetAll(userID)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
	return
}

type DonorRulesWithDeleted struct {
	DonorRules          models.DonorRules `json:"donorRules"`
	DonorRulesForDelete []uuid.UUID       `json:"donorRulesForDelete"`
}

func (h *Handler) UpdateMany(c *gin.Context) {
	var item DonorRulesWithDeleted
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, item.DonorRules, files)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.UpsertMany(item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item.DonorRules)
}

func (h *Handler) AddToUser(c *gin.Context) {
	item := models.DonorRuleUser{}
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	item.UserID, err = h.helper.Token.GetUserID(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
		return
	}
	err = h.service.AddToUser(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) DeleteFromUser(c *gin.Context) {
	item := models.DonorRuleUser{}
	donorRuleID, err := uuid.Parse(c.Param("donor-rule-id"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	item.DonorRuleID = donorRuleID
	item.UserID, err = h.helper.Token.GetUserID(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
		return
	}
	err = h.service.DeleteFromUser(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}
