package hospitalization

import "mdgkb/mdgkb-server/models"

func (s *Service) GetAll() (models.Hospitalizations, error) {
	return s.repository.getAll()
}
