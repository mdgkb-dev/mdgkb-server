package middleware

import (
	"net/http"

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
	return func(_ *gin.Context) {
		// userID, _ := m.helper.Token.GetUserID(c)
		// c.Set("userId", userID)
		//if err != nil {
		//	c.JSON(http.StatusUnauthorized, err)
		//	return
		//}
	}
}

func (m *Middleware) methodIsAllowed(requestMethod string) bool {
	allowedMethods := []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"}
	for _, allowedMethod := range allowedMethods {
		if requestMethod == allowedMethod {
			return true
		}
	}
	return false
}

func (m *Middleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		if !m.methodIsAllowed(c.Request.Method) {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}

		c.Next()
	}
}

func (m *Middleware) CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if !m.checkPermission(c) {
		//	c.AbortWithStatus(http.StatusForbidden)
		//	return
		//}
		c.Next()
	}
}
