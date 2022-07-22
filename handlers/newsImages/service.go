package newsImages

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.NewsImages) error {
	if len(items) == 0 {
		return nil
	}
	err := fileInfos.CreateService(s.helper).UpsertMany(items.GetFileInfos())
	if err != nil {
		return err
	}
	items.SetFileInfoID()
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items models.NewsImages) error {
	if len(items) == 0 {
		return nil
	}
	err := fileInfos.CreateService(s.helper).UpsertMany(items.GetFileInfos())
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
