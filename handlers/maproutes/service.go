package maproutes

import (
	"mdgkb/mdgkb-server/handlers/maproutenodes"
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

func (s *Service) DeleteAll() error {
	return s.repository.DeleteAll()
}

func (s *Service) CreateMany(items models.MapRoutes) error {
	if len(items) == 0 {
		return nil
	}

	err := s.repository.CreateMany(items)
	if err != nil {
		return err
	}

	items.SetIDForChildren()
	maproutenodesService := maproutenodes.CreateService(s.helper)

	err = maproutenodesService.DeleteAll()
	if err != nil {
		return err
	}

	err = maproutenodesService.CreateMany(items.GetMapRouteNodes())
	if err != nil {
		return err
	}

	return nil
}
