package human

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Human) error {
	if item == nil {
		return nil
	}
	return s.repository.create(item)
}

func (s *Service) Update(item *models.Human) error {
	if item == nil {
		return nil
	}
	return s.repository.update(item)
}
