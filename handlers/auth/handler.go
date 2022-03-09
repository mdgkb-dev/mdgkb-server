package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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
	err := h.helper.Email.SendEmail([]string{"lakkinzimusic@gmail.com"}, "Тема", "Тест")
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) RestorePassword(c *gin.Context) {
	//var user *models.User
	//err := c.Bind(&user)
	//if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
	//  return
	//}
	// res, err := h.service.Register(user),
	//if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
	//  return
	//}
	msg := "Тест просодействие"
	from := mail.NewEmail("Просодействие", "pro-assistance@mail.ru")
	subject := fmt.Sprintf("REMINDER -> %s", msg)
	to := mail.NewEmail("lakkinzimusic@gmail.com", "lakkinzimusic@gmail.com")
	message := mail.NewSingleEmail(from, subject, to, msg, msg)
	client := sendgrid.NewSendClient("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")
	res, err := client.Send(message)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, res)
}
