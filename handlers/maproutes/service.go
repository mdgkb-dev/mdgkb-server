package maproutes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetMapRoute(startNodeID string, endNodeID string) (*models.MapRoute, error) {
	return s.repository.GetMapRoute(startNodeID, endNodeID)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
