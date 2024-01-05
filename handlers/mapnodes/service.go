package mapnodes

import (
	"fmt"
	"mdgkb/mdgkb-server/handlers/maproutes"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateMany(items models.MapNodes) error {
	err := s.repository.DeleteAll()
	if err != nil {
		return err
	}

	err = s.repository.CreateMany(items)
	if err != nil {
		return err
	}

	routes := make(models.MapRoutes, 0)
	routes.Calculate(items)
	fmt.Printf("%+v\n", routes)

	maproutesService := maproutes.CreateService(s.helper)

	err = maproutesService.DeleteAll()
	if err != nil {
		return err
	}

	err = maproutesService.CreateMany(routes)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
