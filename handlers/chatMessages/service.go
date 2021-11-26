package chatMessages


import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.ChatMessage) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.ChatMessages, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.ChatMessage, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.ChatMessage) error {
	return s.repository.update(item)
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}
