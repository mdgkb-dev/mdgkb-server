package valuetypes

import "mdgkb/mdgkb-server/models"

func (s *Service) GetAll() (models.ValueTypes, error) {
	return s.repository.getAll()
}
