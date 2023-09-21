package donorrules

import (
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	// id, _ := h.helper.Token.GetUserID(c)
	//if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
	//	return
	//}
	// items, err := h.service.GetAll(id)
	// if h.helper.HTTP.HandleError(c, err) {
	// 	return
	// }
	c.JSON(http.StatusOK, nil)
}

type RulesWithDeleted struct {
	DonorRules          models.DonorRules `json:"donorRules"`
	DonorRulesForDelete []uuid.UUID       `json:"donorRulesForDelete"`
}

func (h *Handler) UpdateMany(c *gin.Context) {
	var item RulesWithDeleted
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.Upload(c, item.DonorRules, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.UpsertMany(item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item.DonorRules)
}

func (h *Handler) AddToUser(c *gin.Context) {
	item := models.DonorRuleUser{}
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	// userID, err := h.helper.Token.GetUserID(c)
	// if h.helper.HTTP.HandleError(c, err) {
	// 	return
	// }
	// item.UserID = *userID
	err = h.service.AddToUser(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) DeleteFromUser(c *gin.Context) {
	item := models.DonorRuleUser{}
	donorRuleID, err := uuid.Parse(c.Param("donor-rule-id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	item.DonorRuleID = donorRuleID
	// userID, err := h.helper.Token.GetUserID(c)
	// if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
	// 	return
	// }
	// item.UserID = *userID
	err = h.service.DeleteFromUser(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}
