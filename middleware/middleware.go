package middleware

import (
	"mdgkb/mdgkb-server/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	helper *helpers.Helper
}

func CreateMiddleware(helper *helpers.Helper) *Middleware {
	return &Middleware{helper: helper}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := m.helper.Token.GetUserID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		return
	}

}
