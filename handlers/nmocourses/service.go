package nmocourses

import (
	"mdgkb/mdgkb-server/handlers/nmocoursespecializations"
	"mdgkb/mdgkb-server/handlers/nmocourseteachers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.NmoCoursesWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get() (*models.NmoCourse, error) {
	item, err := s.repository.get()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.NmoCourse) error {
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.Name, true)
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = nmocourseteachers.CreateService(s.helper).UpsertMany(item.NmoCoursesTeachers)
	if err != nil {
		return err
	}
	err = nmocoursespecializations.CreateService(s.helper).UpsertMany(item.NmoCoursesSpecializations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.NmoCourse) error {
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.Name, true)
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	nmoCourseTeachersService := nmocourseteachers.CreateService(s.helper)
	err = nmoCourseTeachersService.UpsertMany(item.NmoCoursesTeachers)
	if err != nil {
		return err
	}
	err = nmoCourseTeachersService.DeleteMany(item.NmoCoursesTeachersForDelete)
	if err != nil {
		return err
	}
	nmoCourseSpecializationsService := nmocoursespecializations.CreateService(s.helper)
	err = nmoCourseSpecializationsService.UpsertMany(item.NmoCoursesSpecializations)
	if err != nil {
		return err
	}
	err = nmoCourseSpecializationsService.DeleteMany(item.NmoCoursesSpecializationsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.NmoCourses) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
