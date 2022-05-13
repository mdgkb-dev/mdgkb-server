package timetableDays

import (
	"mdgkb/mdgkb-server/handlers/timePeriods"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.TimetableDays) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	err = timePeriods.CreateService(s.repository.getDB()).CreateMany(items.GetTimePeriods())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.TimetableDays) error {
	if len(items) == 0 {
		return nil
	}
	items.SetForeignKeys()
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	timePeriodService := timePeriods.CreateService(s.repository.getDB())
	err = timePeriodService.DeleteMany(items.GetIDForDelete())
	if err != nil {
		return err
	}
	err = timePeriodService.UpsertMany(items.GetTimePeriods())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
