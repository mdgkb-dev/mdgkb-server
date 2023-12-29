package maproutes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetMapRoute(startNodeID string, endNodeID string) (*models.MapRoute, error) {
	return s.repository.GetMapRoute(startNodeID, endNodeID)
}

func (s *Service) UploadMapNodes(items NodesRequest) error {
	// err := s.repository.DeleteAll()
	// if err != nil {
	// 	return err
	// }

	// err = s.repository.UploadMapNodes(items)
	// if err != nil {
	// 	return err
	// }

	routes := make(models.MapRoutes, 0)
	routes.Calculate(items.MapNodes)
	// fmt.Println(routes)

	return nil
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
