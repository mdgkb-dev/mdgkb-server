package residencyApplicationsPointsAchievements

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.ResidencyApplicationPointsAchievements) error {
	if len(items) == 0 {
		return nil
	}
	fileInfosService := fileInfos.CreateService(s.helper)
	err := fileInfosService.UpsertMany(items.GetFileInfos())
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

func (s *Service) UpsertMany(items models.ResidencyApplicationPointsAchievements) error {
	if len(items) == 0 {
		return nil
	}
	fileInfosService := fileInfos.CreateService(s.helper)
	err := fileInfosService.UpsertMany(items.GetFileInfos())
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
