package documentTypes

import (
	"mdgkb/mdgkb-server/handlers/documentTypeFields"
	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/models/schema"
)

func (s *Service) Create(item *models.DocumentType) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = documentTypeFields.CreateService(s.repository.getDB()).CreateMany(item.DocumentTypeFields)
	return err
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
	item.SetIdForChildren()
	DocumentFieldsService := documentTypeFields.CreateService(s.repository.getDB())
	err = DocumentFieldsService.UpsertMany(item.DocumentTypeFields)
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

func (s *Service) GetDocumentsTypesForTablesNames() map[string]string {
	return schema.GetDocumentTypesForTablesNames()
}
