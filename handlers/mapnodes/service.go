package mapnodes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) UploadMapNodes(items models.MapNodes) error {
	err := s.repository.UploadMapNodes(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
