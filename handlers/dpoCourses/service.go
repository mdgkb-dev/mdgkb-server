package dpoCourses

import (
	"mdgkb/mdgkb-server/handlers/dpoCourseDates"
	"mdgkb/mdgkb-server/handlers/dpoCourseSpecializations"
	"mdgkb/mdgkb-server/handlers/dpoCourseTeachers"
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
	item.SetIdForChildren()
	err = dpoCourseTeachers.CreateService(s.helper).UpsertMany(item.DpoCoursesTeachers)
	if err != nil {
		return err
	}
	err = dpoCourseDates.CreateService(s.helper).UpsertMany(item.DpoCoursesDates)
	if err != nil {
		return err
	}
	err = dpoCourseSpecializations.CreateService(s.helper).UpsertMany(item.DpoCoursesSpecializations)
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
	item.SetIdForChildren()
	dpoCourseTeachersService := dpoCourseTeachers.CreateService(s.helper)
	err = dpoCourseTeachersService.UpsertMany(item.DpoCoursesTeachers)
	if err != nil {
		return err
	}
	err = dpoCourseTeachersService.DeleteMany(item.DpoCoursesTeachersForDelete)
	if err != nil {
		return err
	}
	dpoCourseDatesService := dpoCourseDates.CreateService(s.helper)
	err = dpoCourseDatesService.UpsertMany(item.DpoCoursesDates)
	if err != nil {
		return err
	}
	err = dpoCourseDatesService.DeleteMany(item.DpoCoursesDatesForDelete)
	if err != nil {
		return err
	}
	dpoCourseSpecializationsService := dpoCourseSpecializations.CreateService(s.helper)
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
