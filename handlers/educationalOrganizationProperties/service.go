package educationalOrganizationProperties

import "mdgkb/mdgkb-server/models"

func (s *Service) GetAll() (models.EducationalOrganizationProperties, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpsertMany(items models.EducationalOrganizationProperties) error {
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(id []string) error {
	if len(id) == 0 {
		return nil
	}
	return s.repository.deleteMany(id)
}
