package admissioncommitteedocumenttypes

import (
	"mdgkb/mdgkb-server/handlers/pagesections"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) GetAll() (models.AdmissionCommitteeDocumentTypes, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.AdmissionCommitteeDocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.AdmissionCommitteeDocumentType) error {
	err := pagesections.CreateService(s.helper).Create(item.DocumentType)
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

func (s *Service) Update(item *models.AdmissionCommitteeDocumentType) error {
	err := pagesections.CreateService(s.helper).Update(item.DocumentType)
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

func (s *Service) UpdateOrder(items models.AdmissionCommitteeDocumentTypes) error {
	return s.repository.upsertMany(items)
}

func (s *Service) UpsertMany(items CommitteeDocumentTypesWithDelete) error {
	documentService := pagesections.CreateService(s.helper)
	if len(items.AdmissionCommitteeDocumentTypes) > 0 {
		err := documentService.UpsertMany(items.AdmissionCommitteeDocumentTypes.GetDocumentTypes())
		if err != nil {
			return err
		}
	}
	if len(items.AdmissionCommitteeDocumentTypesForDelete) > 0 {
		items.AdmissionCommitteeDocumentTypes.SetForeignKeys()
		err := s.repository.deleteMany(items.AdmissionCommitteeDocumentTypesForDelete)
		if err != nil {
			return err
		}
	}
	items.AdmissionCommitteeDocumentTypes.SetForeignKeys()
	if len(items.AdmissionCommitteeDocumentTypes) > 0 {
		err := s.repository.upsertMany(items.AdmissionCommitteeDocumentTypes)
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
