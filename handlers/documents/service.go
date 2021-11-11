package documents

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/documentsScans"
	"mdgkb/mdgkb-server/models"
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
	err = documentsScans.CreateService(s.repository.getDB()).UpsertMany(items.GetDocumentsScans())
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
	err = documentsScans.CreateService(s.repository.getDB()).UpsertMany(items.GetDocumentsScans())
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
