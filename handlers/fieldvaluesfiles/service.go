package fieldvaluesfiles

import (
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.FieldValuesFiles) error {
	if len(items) == 0 {
		return nil
	}
	err := fileinfos.CreateService(s.helper).UpsertMany(items.GetFileInfos())
	if err != nil {
		return err
	}
	items.SetFileInfoID()
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items models.FieldValuesFiles) error {
	if len(items) == 0 {
		return nil
	}
	err := fileinfos.CreateService(s.helper).UpsertMany(items.GetFileInfos())
	if err != nil {
		return err
	}
	items.SetFileInfoID()
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
