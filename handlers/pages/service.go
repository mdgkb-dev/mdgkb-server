package pages

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Page) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.Pages, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.Page, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Page) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) GetBySlug(slug *string) (*models.Page, error) {
	item, err := s.repository.getBySlug(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}