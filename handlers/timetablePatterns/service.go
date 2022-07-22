package timetablePatterns

import (
	"mdgkb/mdgkb-server/handlers/timetableDays"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.TimetablePatterns, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.TimetablePattern, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.TimetablePattern) error {
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

func (s *Service) Update(item *models.TimetablePattern) error {
	err := s.repository.update(item)
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

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}
