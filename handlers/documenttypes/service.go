package documenttypes

import (
	"mdgkb/mdgkb-server/handlers/documents"
	"mdgkb/mdgkb-server/handlers/documenttypefields"
	"mdgkb/mdgkb-server/handlers/documenttypesimages"
	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/models/schema"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.DocumentType) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = documenttypefields.CreateService(s.helper).CreateMany(item.DocumentTypeFields)
	if err != nil {
		return err
	}
	err = documents.CreateService(s.helper).UpsertMany(item.Documents)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(params models.DocumentsParams) ([]*models.DocumentType, error) {
	items, err := s.repository.getAll(params)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.DocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.DocumentType) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	documentFieldsService := documenttypefields.CreateService(s.helper)
	err = documentFieldsService.UpsertMany(item.DocumentTypeFields)
	if err != nil {
		return err
	}
	if len(item.DocumentFieldsForDelete) > 0 {
		err = documentFieldsService.DeleteMany(item.DocumentFieldsForDelete)
		if err != nil {
			return err
		}
	}

	documentService := documents.CreateService(s.helper)
	if len(item.DocumentsForDelete) > 0 {
		err = documentService.DeleteMany(item.DocumentsForDelete)
		if err != nil {
			return err
		}
	}
	err = documentService.UpsertMany(item.Documents)
	if err != nil {
		return err
	}

	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) GetDocumentsTypesForTablesNames() map[string]string {
	return schema.GetDocumentTypesForTablesNames()
}

func (s *Service) Upsert(item *models.DocumentType) error {
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	documentService := documents.CreateService(s.helper)
	err = documentService.DeleteMany(item.DocumentsForDelete)
	if err != nil {
		return err
	}
	err = documentService.UpsertMany(item.Documents)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.DocumentTypes) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	documentService := documents.CreateService(s.helper)
	err = documentService.DeleteMany(items.GetDocumentsIDForDelete())
	if err != nil {
		return err
	}
	err = documentService.UpsertMany(items.GetDocuments())
	if err != nil {
		return err
	}
	documentFieldsService := documenttypesimages.CreateService(s.helper)
	err = documentFieldsService.UpsertMany(items.GetDocumentTypeImages())
	if err != nil {
		return err
	}
	err = documentFieldsService.DeleteMany(items.GetDocumentTypeImagesForDelete())
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
