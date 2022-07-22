package candidateDocumentTypes

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/documentTypes"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.CandidateDocumentTypes, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.CandidateDocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.CandidateDocumentType) error {
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

func (s *Service) Update(item *models.CandidateDocumentType) error {
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

func (s *Service) UpsertMany(items CandidateDocumentTypesWithDelete) error {
	documentService := documentTypes.CreateService(s.helper)
	if len(items.CandidateDocumentTypes) > 0 {
		err := documentService.UpsertMany(items.CandidateDocumentTypes.GetDocumentTypes())
		if err != nil {
			return err
		}
	}
	if len(items.CandidateDocumentTypesForDelete) > 0 {
		items.CandidateDocumentTypes.SetForeignKeys()
		err := s.repository.deleteMany(items.CandidateDocumentTypesForDelete)
		if err != nil {
			return err
		}
	}
	items.CandidateDocumentTypes.SetForeignKeys()
	if len(items.CandidateDocumentTypes) > 0 {
		err := s.repository.upsertMany(items.CandidateDocumentTypes)
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
