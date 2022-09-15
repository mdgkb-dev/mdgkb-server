package divisions

import (
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/contactinfo"
	"mdgkb/mdgkb-server/handlers/divisionimages"
	"mdgkb/mdgkb-server/handlers/divisionvideos"
	"mdgkb/mdgkb-server/handlers/doctors"
	"mdgkb/mdgkb-server/handlers/schedules"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/handlers/visitingrules"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Division) error {
	timetableService := timetables.CreateService(s.helper)
	err := timetableService.Create(item.Timetable)
	if err != nil {
		return err
	}
	schedulesService := schedules.CreateService(s.helper)
	err = schedulesService.Create(item.Schedule)
	if err != nil {
		return err
	}
	item.Slug = s.helper.Util.MakeSlug(item.Name, true)

	contactInfoService := contactinfo.CreateService(s.helper)
	err = contactInfoService.Create(item.ContactInfo)
	if err != nil {
		return err
	}
	item.SetForeignKeys()

	doctorsService := doctors.CreateService(s.helper)
	err = doctorsService.UpsertMany(item.Doctors)
	if err != nil {
		return err
	}

	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	visitingRulesService := visitingrules.CreateService(s.helper)
	err = visitingRulesService.UpsertMany(item.VisitingRules)
	if err != nil {
		return err
	}
	divisionImagesService := divisionimages.CreateService(s.helper)
	err = divisionImagesService.CreateMany(item.DivisionImages)
	if err != nil {
		return err
	}
	divisionVideosService := divisionvideos.CreateService(s.helper)
	err = divisionVideosService.CreateMany(item.DivisionVideos)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Division) error {
	timetableService := timetables.CreateService(s.helper)
	err := timetableService.Upsert(item.Timetable)
	if err != nil {
		return err
	}

	schedulesService := schedules.CreateService(s.helper)
	err = schedulesService.Upsert(item.Schedule)
	if err != nil {
		return err
	}

	contactInfoService := contactinfo.CreateService(s.helper)
	err = contactInfoService.Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	item.SetForeignKeys()

	doctorsService := doctors.CreateService(s.helper)
	err = doctorsService.DeleteMany(item.DoctorsForDelete)
	if err != nil {
		return err
	}
	err = doctorsService.UpsertMany(item.Doctors)
	if err != nil {
		return err
	}

	divisionImagesService := divisionimages.CreateService(s.helper)
	err = divisionImagesService.DeleteMany(item.DivisionImagesForDelete)
	if err != nil {
		return err
	}
	err = divisionImagesService.UpsertMany(item.DivisionImages)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	visitingRulesService := visitingrules.CreateService(s.helper)
	err = visitingRulesService.UpsertMany(item.VisitingRules)
	if err != nil {
		return err
	}
	err = visitingRulesService.DeleteMany(item.VisitingRulesForDelete)
	if err != nil {
		return err
	}
	divisionVideosService := divisionvideos.CreateService(s.helper)
	err = divisionVideosService.UpsertMany(item.DivisionVideos)
	if err != nil {
		return err
	}
	err = divisionVideosService.DeleteMany(item.DivisionVideosForDelete)
	if err != nil {
		return err
	}
	return s.repository.update(item)
}

func (s *Service) GetAll() (models.DivisionsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get() (*models.Division, error) {
	item, err := s.repository.get()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) CreateComment(item *models.DivisionComment) error {
	commentsService := comments.CreateService(s.helper)
	err := commentsService.UpsertOne(item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return s.repository.createComment(item)
}

func (s *Service) UpdateComment(item *models.DivisionComment) error {
	commentsService := comments.CreateService(s.helper)
	err := commentsService.UpdateOne(item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
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
