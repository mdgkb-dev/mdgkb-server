package divisions

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/handlers/divisionImages"
	"mdgkb/mdgkb-server/handlers/doctors"
	"mdgkb/mdgkb-server/handlers/schedules"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/handlers/visitingRules"
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
	item.Slug = s.helper.MakeSlug(item.Name)

	doctorsService := doctors.CreateService(s.repository.getDB(), s.helper)
	err = doctorsService.UpsertMany(item.Doctors)
	if err != nil {
		return err
	}

	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	visitingRulesService := visitingRules.CreateService(s.repository.getDB(), s.helper)
	err = visitingRulesService.UpsertMany(item.VisitingRules)
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

	doctorsService := doctors.CreateService(s.repository.getDB(), s.helper)
	err = doctorsService.UpsertMany(item.Doctors)
	if err != nil {
		return err
	}

	divisionImagesService := divisionImages.CreateService(s.repository.getDB())
	err = divisionImagesService.DeleteMany(item.DivisionImagesForDelete)
	if err != nil {
		return err
	}
	err = divisionImagesService.UpsertMany(item.DivisionImages)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	visitingRulesService := visitingRules.CreateService(s.repository.getDB(), s.helper)
	err = visitingRulesService.UpsertMany(item.VisitingRules)
	if err != nil {
		return err
	}
	err = visitingRulesService.DeleteMany(item.VisitingRulesForDelete)
	if err != nil {
		return err
	}
	return s.repository.update(item)
}

func (s *Service) GetAll(onlyShowed bool) (models.Divisions, error) {
	return s.repository.getAll(onlyShowed)
}

func (s *Service) Get(id string, onlyShowed bool) (*models.Division, error) {
	item, err := s.repository.get(id, onlyShowed)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) CreateComment(item *models.DivisionComment) error {
	return s.repository.createComment(item)
}

func (s *Service) UpdateComment(item *models.DivisionComment) error {
	return s.repository.updateComment(item)
}

func (s *Service) RemoveComment(id string) error {
	return s.repository.removeComment(id)
}

func (s *Service) GetBySearch(search string) (models.Divisions, error) {
	return s.repository.getBySearch(search)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
