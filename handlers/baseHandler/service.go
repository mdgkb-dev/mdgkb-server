package baseHandler

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
