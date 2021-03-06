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
	restoreLink := h.helper.HTTP.GetRestorePasswordURL(findedUser.ID.String(), findedUser.UUID.String())
	err = h.helper.Email.SendEmail([]string{user.Email}, "???????????????????????????? ???????????? ?????? ?????????????? ??????????", restoreLink)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, err)
}

func (h *Handler) CheckUUID(c *gin.Context) {
	findedUser, err := h.service.GetUserByID(c.Param("user-id"))
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

func (h *Handler) SavePathPermissions(c *gin.Context) {
	var items models.PathPermissions
	err := c.Bind(&items)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.UpsertManyPathPermissions(items)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, err)
}

func (h *Handler) GetAllPathPermissions(c *gin.Context) {
	items, err := h.service.GetAllPathPermissions()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetAllPathPermissionsAdmin(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	items, err := h.service.GetAllPathPermissionsAdmin()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetPathPermissionsByRoleId(c *gin.Context) {
	items, err := h.service.GetPathPermissionsByRoleId(c.Param("roleId"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) CheckPathPermissions(c *gin.Context) {
	var path string
	err := c.Bind(&path)
	if h.helper.HTTP.HandleError(c, err, http.StatusForbidden) {
		return
	}
	userRoleId := ""
	if c.Request.Header.Get("token") != "null" {
		accessDetails, err := h.helper.Token.GetAccessDetail(c)
		if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
			return
		}
		userRoleId = accessDetails.UserRoleID
	}
	err = h.service.CheckPathPermissions(path, userRoleId)
	if h.helper.HTTP.HandleError(c, err, http.StatusForbidden) {
		return
	}
	c.JSON(http.StatusOK, nil)
}
