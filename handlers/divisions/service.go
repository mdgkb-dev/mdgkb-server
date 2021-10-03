package divisions

import (
	"mdgkb/mdgkb-server/handlers/divisionImages"
	"mdgkb/mdgkb-server/handlers/schedules"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Division) error {
	timetableService := timetables.CreateService(s.repository.getDB())
	err := timetableService.Create(item.Timetable)
	item.TimetableId = item.Timetable.ID
	if err != nil {
		return err
	}
	schedulesService := schedules.CreateService(s.repository.getDB())
	err = schedulesService.Create(item.Schedule)
	if err != nil {
		return err
	}
	item.ScheduleId = item.Schedule.ID

	err = s.repository.create(item)
	if err != nil {
		return err
	}

	divisionImagesService := divisionImages.CreateService(s.repository.getDB())
	err = divisionImagesService.CreateMany(item.DivisionImages)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Division) error {
	timetableService := timetables.CreateService(s.repository.getDB())
	err := timetableService.Upsert(item.Timetable)
	if err != nil {
		return err
	}
	item.TimetableId = item.Timetable.ID

	schedulesService := schedules.CreateService(s.repository.getDB())
	err = schedulesService.Upsert(item.Schedule)
	if err != nil {
		return err
	}
	item.ScheduleId = item.Schedule.ID


	divisionImagesService := divisionImages.CreateService(s.repository.getDB())
	err = divisionImagesService.DeleteMany(item.DivisionImagesForDelete)
	if err != nil {
		return err
	}
	err = divisionImagesService.UpsertMany(item.DivisionImages)
	if err != nil {
		return err
	}

	return s.repository.update(item)
}


func (s *Service) GetAll() (models.Divisions, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.Division, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) CreateComment(item *models.DivisionComment) error {
	return s.repository.createComment(item)
}

func (s *Service) UpdateComment(item *models.DivisionComment) error {
	return s.repository.updateComment(item)
}

func (s *Service) RemoveComment(id *string) error {
	return s.repository.removeComment(id)
}