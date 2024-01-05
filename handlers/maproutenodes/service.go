package maproutenodes

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) DeleteAll() error {
	return s.repository.DeleteAll()
}

func (s *Service) CreateMany(items models.MapRouteNodes) error {
	if len(items) == 0 {
		return nil
	}

	err := s.repository.CreateMany(items)
	if err != nil {
		return err
	}

	return nil
}
