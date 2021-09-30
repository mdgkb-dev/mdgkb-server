package educationalOrganizationPages

import (
	"mdgkb/mdgkb-server/handlers/pages"
	"mdgkb/mdgkb-server/models"
)



func (s *Service) CreateMany(items models.EducationalOrganizationPages) error {
	if len(items) == 0 {
		return nil
	}
	err := pages.CreateService(s.repository.getDB()).CreateMany(items.GetPages())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.createMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.EducationalOrganizationPages, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpsertMany(items models.EducationalOrganizationPages) error {
	if len(items) == 0 {
		return nil
	}
	err := pages.CreateService(s.repository.getDB()).UpsertMany(items.GetPages())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
