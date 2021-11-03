package document

import (
	"mdgkb/mdgkb-server/handlers/documentFields"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Document) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = documentFields.CreateService(s.repository.getDB()).CreateMany(item.DocumentFields)
	return err
}

func (s *Service) GetAll() ([]*models.Document, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.Document, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Document) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	DocumentFieldsService := documentFields.CreateService(s.repository.getDB())
	err = DocumentFieldsService.UpsertMany(item.DocumentFields)
	if err != nil {
		return err
	}
	if len(item.DocumentFieldsForDelete) > 0 {
		err = DocumentFieldsService.DeleteMany(item.DocumentFieldsForDelete)
	}
	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
