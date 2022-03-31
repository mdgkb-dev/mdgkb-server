package dpoDocumentTypes

import (
	"fmt"
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/documentTypes"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.DpoDocumentTypes, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.DpoDocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.DpoDocumentType) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	//item.SetIdForChildren()
	//err = documentTypes.CreateService(s.repository.getDB(), s.helper).UpsertMany(item.DocumentTypes)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s *Service) Update(item *models.DpoDocumentType) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	//item.SetIdForChildren()
	//documentTypeService := documentTypes.CreateService(s.repository.getDB(), s.helper)
	//err = documentTypeService.DeleteMany(item.DocumentTypesForDelete)
	//if err != nil {
	//	return err
	//}
	//err = documentTypeService.UpsertMany(item.DocumentTypes)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(items DpoDocumentTypesWithDelete) error {
	fmt.Println(items.DpoDocumentTypes)
	documentService := documentTypes.CreateService(s.repository.getDB(), s.helper)
	if len(items.DpoDocumentTypes) > 0 {
		err := documentService.UpsertMany(items.DpoDocumentTypes.GetDocumentTypes())
		if err != nil {
			return err
		}
	}
	if len(items.DpoDocumentTypesForDelete) > 0 {
		items.DpoDocumentTypes.SetForeignKeys()
		err := s.repository.deleteMany(items.DpoDocumentTypesForDelete)
		if err != nil {
			return err
		}
	}
	items.DpoDocumentTypes.SetForeignKeys()
	if len(items.DpoDocumentTypes) > 0 {
		err := s.repository.upsertMany(items.DpoDocumentTypes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
