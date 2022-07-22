package postgraduateDocumentTypes

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/documentTypes"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.PostgraduateDocumentTypes, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.PostgraduateDocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.PostgraduateDocumentType) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	//item.SetIdForChildren()
	//err = documentTypes.CreateService(s.helper).UpsertMany(item.DocumentTypes)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s *Service) Update(item *models.PostgraduateDocumentType) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	//item.SetIdForChildren()
	//documentTypeService := documentTypes.CreateService(s.helper)
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

func (s *Service) UpsertMany(items PostgraduateDocumentTypesWithDelete) error {
	documentService := documentTypes.CreateService(s.helper)
	if len(items.PostgraduateDocumentTypes) > 0 {
		err := documentService.UpsertMany(items.PostgraduateDocumentTypes.GetDocumentTypes())
		if err != nil {
			return err
		}
	}
	if len(items.PostgraduateDocumentTypesForDelete) > 0 {
		items.PostgraduateDocumentTypes.SetForeignKeys()
		err := s.repository.deleteMany(items.PostgraduateDocumentTypesForDelete)
		if err != nil {
			return err
		}
	}
	items.PostgraduateDocumentTypes.SetForeignKeys()
	if len(items.PostgraduateDocumentTypes) > 0 {
		err := s.repository.upsertMany(items.PostgraduateDocumentTypes)
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
