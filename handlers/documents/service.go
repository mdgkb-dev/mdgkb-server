package documents

import (
	"mdgkb/mdgkb-server/handlers/documentsScans"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.Documents) error {
	if len(items) == 0 {
		return nil
	}

	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	err = documentsScans.CreateService(s.helper).UpsertMany(items.GetDocumentsScans())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.Documents) error {
	if len(items) == 0 {
		return nil
	}

	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	documentsScanService := documentsScans.CreateService(s.helper)
	err = documentsScanService.DeleteMany(items.GetDocumentsScansIdForDelete())
	if err != nil {
		return err
	}
	err = documentsScanService.UpsertMany(items.GetDocumentsScans())
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
