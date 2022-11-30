package pagesectiondocuments

import (
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.PageSectionDocuments) error {
	if len(items) == 0 {
		return nil
	}
	err := fileinfos.CreateService(s.helper).UpsertMany(items.GetScans())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.createMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.PageSectionDocuments) error {
	if len(items) == 0 {
		return nil
	}

	err := fileinfos.CreateService(s.helper).UpsertMany(items.GetScans())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.upsertMany(items)
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
