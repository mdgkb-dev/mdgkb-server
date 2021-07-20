package middleware

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/helpers"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Extract the access token metadata
		_, err := helpers.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
	}
}
