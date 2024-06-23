package auth

import (
	"fmt"
	"net/http"

	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	baseModels "github.com/pro-assistance/pro-assister/models"
)

func (h *Handler) Register(c *gin.Context) {
	var item *models.User
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	res, err := S.Register(c.Request.Context(), item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var item baseModels.UserAccount
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	res, err := S.Login(c.Request.Context(), item.Email, item.Password)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, res)
}

// func (h *Handler) LoginAs(c *gin.Context) {
// 	var item models.Login
// 	err := c.Bind(&item)
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	res, err := h.service.Login(&item, true)
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	c.JSON(http.StatusOK, res)
// }

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

type refreshToken struct {
	RefreshToken string `json:"refreshToken"`
	UserID       string `json:"userId"`
}

//	func (h *Handler) RefreshToken(c *gin.Context) {
//		t := refreshToken{}
//		err := c.Bind(&t)
//		if h.helper.HTTP.HandleError(c, err) {
//			return
//		}
//		//userId, err := h.helper.Token.ExtractTokenMetadata(c.Request, "user_id")
//		//if h.helper.HTTP.HandleError(c, err) {
//		//	return
//		//}
//		user, err := users.CreateService(h.helper).Get(t.UserID)
//		if h.helper.HTTP.HandleError(c, err) {
//			return
//		}
//		tokens, err := h.helper.Token.RefreshToken(t.RefreshToken, user)
//		if h.helper.HTTP.HandleError(c, err) {
//			return
//		}
//		c.JSON(http.StatusOK, tokens)
//	}
func (h *Handler) RefreshPassword(c *gin.Context) {
	var item models.User
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.UpdatePassword(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) RestorePassword(c *gin.Context) {
	var item baseModels.UserAccount
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.RestorePassword(c, item.Email)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, err)
}

//
// func (h *Handler) SavePathPermissions(c *gin.Context) {
// 	var items models.PathPermissions
// 	err := c.Bind(&items)
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	err = h.service.UpsertManyPathPermissions(items)
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	c.JSON(http.StatusOK, err)
// }
//
// func (h *Handler) GetAllPathPermissions(c *gin.Context) {
// 	items, err := h.service.GetAllPathPermissions()
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	c.JSON(http.StatusOK, items)
// }
//
// func (h *Handler) GetAllPathPermissionsAdmin(c *gin.Context) {
// 	err := h.service.setQueryFilter(c)
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	items, err := h.service.GetAllPathPermissionsAdmin()
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	c.JSON(http.StatusOK, items)
// }
//
// func (h *Handler) GetPathPermissionsByRoleID(c *gin.Context) {
// 	items, err := h.service.GetPathPermissionsByRoleID(c.Param("roleId"))
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	c.JSON(http.StatusOK, items)
// }
//
// func (h *Handler) CheckPathPermissions(c *gin.Context) {
// 	var path string
// 	err := c.Bind(&path)
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	userRoleID := ""
// 	err = h.service.CheckPathPermissions(path, userRoleID)
// 	if h.helper.HTTP.HandleError(c, err) {
// 		return
// 	}
// 	c.JSON(http.StatusOK, nil)
// }
