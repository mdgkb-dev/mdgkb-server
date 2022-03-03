package middleware

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/helpers"
)

type Middleware struct {
	helper *helpers.Helper
}

func CreateMiddleware(helper *helpers.Helper) *Middleware {
	return &Middleware{helper: helper}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := m.helper.Token.GetUserID(c)
		c.Set("userId", userID)
		//if err != nil {
		//	c.JSON(http.StatusUnauthorized, err)
		//	return
		//}
		return
	}

}
