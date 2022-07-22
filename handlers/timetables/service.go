package timetables

import (
	"mdgkb/mdgkb-server/handlers/timetableDays"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Timetable) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	timetableDaysService := timetableDays.CreateService(s.helper)
	err = timetableDaysService.CreateMany(item.TimetableDays)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Upsert(item *models.Timetable) error {
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	timetableDaysService := timetableDays.CreateService(s.helper)
	err = timetableDaysService.UpsertMany(item.TimetableDays)
	if err != nil {
		return err
	}
	err = timetableDaysService.DeleteMany(item.TimetableDaysForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAllWeekdays() (models.Weekdays, error) {
	items, err := s.repository.getAllWeekdays()
	if err != nil {
		return nil, err
	}
	return items, nil
}
