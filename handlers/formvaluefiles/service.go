package formvaluefiles

import (
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.FormValueFiles) error {
	if len(items) == 0 {
		return nil
	}
	err := fileinfos.CreateService(s.helper).UpsertMany(items.GetFiles())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items models.FormValueFiles) error {
	if len(items) == 0 {
		return nil
	}
	err := fileinfos.CreateService(s.helper).UpsertMany(items.GetFiles())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	return s.repository.upsertMany(items)
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
