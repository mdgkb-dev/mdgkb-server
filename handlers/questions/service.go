package questions

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Question) error {

	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(published bool) (models.Questions, error) {
	items, err := s.repository.getAll(published)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.Question, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Question) error {
	return s.repository.update(item)
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) ReadAnswers(userID string) error {
	return s.repository.readAnswers(userID)
}

func (s *Service) Publish(id string) error {
	return s.repository.publish(id)
}
