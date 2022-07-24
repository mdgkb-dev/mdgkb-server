package dpocourses

import (
	"mdgkb/mdgkb-server/handlers/dpocoursedates"
	"mdgkb/mdgkb-server/handlers/dpocoursespecializations"
	"mdgkb/mdgkb-server/handlers/dpocourseteachers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.DpoCoursesWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get() (*models.DpoCourse, error) {
	item, err := s.repository.get()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.DpoCourse) error {
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.Name, true)
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = dpocourseteachers.CreateService(s.helper).UpsertMany(item.DpoCoursesTeachers)
	if err != nil {
		return err
	}
	err = dpocoursedates.CreateService(s.helper).UpsertMany(item.DpoCoursesDates)
	if err != nil {
		return err
	}
	err = dpocoursespecializations.CreateService(s.helper).UpsertMany(item.DpoCoursesSpecializations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.DpoCourse) error {
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.Name, true)
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	dpoCourseTeachersService := dpocourseteachers.CreateService(s.helper)
	err = dpoCourseTeachersService.UpsertMany(item.DpoCoursesTeachers)
	if err != nil {
		return err
	}
	err = dpoCourseTeachersService.DeleteMany(item.DpoCoursesTeachersForDelete)
	if err != nil {
		return err
	}
	dpoCourseDatesService := dpocoursedates.CreateService(s.helper)
	err = dpoCourseDatesService.UpsertMany(item.DpoCoursesDates)
	if err != nil {
		return err
	}
	err = dpoCourseDatesService.DeleteMany(item.DpoCoursesDatesForDelete)
	if err != nil {
		return err
	}
	dpoCourseSpecializationsService := dpocoursespecializations.CreateService(s.helper)
	err = dpoCourseSpecializationsService.UpsertMany(item.DpoCoursesSpecializations)
	if err != nil {
		return err
	}
	err = dpoCourseSpecializationsService.DeleteMany(item.DpoCoursesSpecializationsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.DpoCourses) error {
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
