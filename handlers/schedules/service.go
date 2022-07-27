package schedules

import (
	"mdgkb/mdgkb-server/handlers/scheduleitems"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Schedule) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	scheduleItemsService := scheduleitems.CreateService(s.helper)
	err = scheduleItemsService.CreateMany(item.ScheduleItems)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Upsert(item *models.Schedule) error {
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	scheduleItemsService := scheduleitems.CreateService(s.helper)
	err = scheduleItemsService.UpsertMany(item.ScheduleItems)
	if err != nil {
		return err
	}
	err = scheduleItemsService.DeleteMany(item.ScheduleItemsForDelete)
	if err != nil {
		return err
	}
	return s.repository.upsert(item)
}
