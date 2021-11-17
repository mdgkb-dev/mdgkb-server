package hospitalization

import "mdgkb/mdgkb-server/models"

func (s *Service) GetAll() (models.Hospitalizations, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.Hospitalization, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
