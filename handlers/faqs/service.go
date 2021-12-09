package faqs

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Faq) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.Faqs, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.Faq, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) UpsertMany(items models.FaqsWithDelete) error {
	err := s.repository.upsertMany(items.Faqs)
	if err != nil {
		return err
	}
	if len(items.FaqsForDelete) > 0 {
		err = s.repository.deleteMany(items.FaqsForDelete)
	}
	return err
}

func (s *Service) Update(item *models.Faq) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}
