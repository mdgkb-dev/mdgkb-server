package httpHelper

import "github.com/gin-gonic/gin"

func GetID(c *gin.Context) *string {
	id := c.Param("id")
	return &id
}