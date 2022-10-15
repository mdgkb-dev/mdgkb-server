package educationdocumenttypes

import (
	"mdgkb/mdgkb-server/handlers/documenttypes"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) GetAll() (models.EducationDocumentTypes, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.EducationDocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.EducationDocumentType) error {
	err := documenttypes.CreateService(s.helper).Create(item.DocumentType)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.EducationDocumentType) error {
	err := documenttypes.CreateService(s.helper).Update(item.DocumentType)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) UpdateOrder(items models.EducationDocumentTypes) error {
	return s.repository.upsertMany(items)
}

func (s *Service) UpsertMany(items CommitteeDocumentTypesWithDelete) error {
	documentService := documenttypes.CreateService(s.helper)
	if len(items.EducationDocumentTypes) > 0 {
		err := documentService.UpsertMany(items.EducationDocumentTypes.GetDocumentTypes())
		if err != nil {
			return err
		}
	}
	if len(items.EducationDocumentTypesForDelete) > 0 {
		items.EducationDocumentTypes.SetForeignKeys()
		err := s.repository.deleteMany(items.EducationDocumentTypesForDelete)
		if err != nil {
			return err
		}
	}
	items.EducationDocumentTypes.SetForeignKeys()
	if len(items.EducationDocumentTypes) > 0 {
		err := s.repository.upsertMany(items.EducationDocumentTypes)
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
