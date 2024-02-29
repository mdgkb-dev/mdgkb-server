package pagesections

import (
	"fmt"
	"mdgkb/mdgkb-server/handlers/pagesectiondocuments"
	"mdgkb/mdgkb-server/handlers/pagesectionimages"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.PageSection) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = pagesectiondocuments.CreateService(s.helper).UpsertMany(item.PageSectionDocuments)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(params models.DocumentsParams) ([]*models.PageSection, error) {
	items, err := s.repository.getAll(params)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.PageSection, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.PageSection) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	documentService := pagesectiondocuments.CreateService(s.helper)
	if len(item.PageSectionDocumentsForDelete) > 0 {
		err = documentService.DeleteMany(item.PageSectionDocumentsForDelete)
		if err != nil {
			return err
		}
	}
	err = documentService.UpsertMany(item.PageSectionDocuments)
	if err != nil {
		return err
	}

	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) Upsert(item *models.PageSection) error {
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	documentService := pagesectiondocuments.CreateService(s.helper)
	err = documentService.DeleteMany(item.PageSectionDocumentsForDelete)
	if err != nil {
		return err
	}
	err = documentService.UpsertMany(item.PageSectionDocuments)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.PageSections) error {
	if len(items) == 0 {
		return nil
	}
	fmt.Println("2.4.1")
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	fmt.Println("2.4.2")
	documentService := pagesectiondocuments.CreateService(s.helper)
	err = documentService.UpsertMany(items.GetDocuments())
	if err != nil {
		return err
	}
	err = documentService.DeleteMany(items.GetDocumentsIDForDelete())
	if err != nil {
		return err
	}
	fmt.Println("2.4.3")
	documentFieldsService := pagesectionimages.CreateService(s.helper)
	err = documentFieldsService.UpsertMany(items.GetDocumentTypeImages())
	if err != nil {
		return err
	}
	fmt.Println("2.4.3")
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
