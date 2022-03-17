package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
)

type Middleware struct {
	helper *helper.Helper
}

func CreateMiddleware(helper *helper.Helper) *Middleware {
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
