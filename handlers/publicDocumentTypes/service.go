package publicDocumentTypes

import (
	"mdgkb/mdgkb-server/handlers/documentTypes"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.PublicDocumentTypes, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.PublicDocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.PublicDocumentType) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = documentTypes.CreateService(s.repository.getDB()).UpsertMany(item.DocumentTypes)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.PublicDocumentType) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	documentTypeService := documentTypes.CreateService(s.repository.getDB())
	err = documentTypeService.DeleteMany(item.DocumentTypesForDelete)
	if err != nil {
		return err
	}
	err = documentTypeService.UpsertMany(item.DocumentTypes)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}
