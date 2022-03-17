package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {
	var user *models.User
	err := c.Bind(&user)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = user.GenerateHashPassword()

	item, err := h.service.Register(user)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Login(c *gin.Context) {
	var item models.User
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	res, err := h.service.Login(&item)

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	//_, err := h.helper.Token.ExtractTokenMetadata(c.Request)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, "unauthorized")
	//	return
	//}
	//delErr := helpers.DeleteTokens(metadata, h.redis)
	//if delErr != nil {
	//	c.JSON(http.StatusUnauthorized, delErr.Error())
	//	return
	//}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func (h *Handler) RefreshToken(c *gin.Context) {
	type refreshToken struct {
		RefreshToken string `json:"refreshToken"`
	}
	t := refreshToken{}
	err := c.Bind(&t)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	tokens, err := h.helper.Token.RefreshToken(t.RefreshToken)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) RefreshPassword(c *gin.Context) {
	var item models.User
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.UpdatePassword(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) RestorePassword(c *gin.Context) {
	var user *models.User
	err := c.Bind(&user)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	findedUser, err := h.service.FindUserByEmail(user.Email)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	restoreLink, err := h.helper.HTTP.GetRestorePasswordURL(findedUser.ID.String(), findedUser.UUID.String())
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.helper.Email.SendEmail([]string{user.Email}, "Восстановление пароля для портала МДГКБ", restoreLink)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, err)
}

func (h *Handler) checkUUID(c *gin.Context) {
	findedUser, err := h.service.GetUserByID(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	if !findedUser.CompareWithUUID(c.Param("uuid")) {
		err = errors.New("wrong unique signature")
		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
	}
	err = h.service.DropUUID(findedUser)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, nil)
}
