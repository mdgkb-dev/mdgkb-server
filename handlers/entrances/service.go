package entrances

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.Entrances, error) {
	return s.repository.getAll()
}
