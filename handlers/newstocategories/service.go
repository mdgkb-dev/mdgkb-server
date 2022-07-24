package newstocategories

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.NewsToCategories) error {
	if len(items) == 0 {
		return nil
	}

	err := s.repository.createMany(items)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpsertMany(items models.NewsToCategories) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
