package dpoCourses

import (
	"mdgkb/mdgkb-server/handlers/dpoCourseDates"
	"mdgkb/mdgkb-server/handlers/dpoCourseSpecializations"
	"mdgkb/mdgkb-server/handlers/dpoCourseTeachers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.DpoCourses, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.DpoCourse, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.DpoCourse) error {
	item.SetForeignKeys()
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = dpoCourseTeachers.CreateService(s.repository.getDB()).UpsertMany(item.DpoCoursesTeachers)
	if err != nil {
		return err
	}
	err = dpoCourseDates.CreateService(s.repository.getDB()).UpsertMany(item.DpoCoursesDates)
	if err != nil {
		return err
	}
	err = dpoCourseSpecializations.CreateService(s.repository.getDB()).UpsertMany(item.DpoCoursesSpecializations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.DpoCourse) error {
	item.SetForeignKeys()
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	dpoCourseTeachersService := dpoCourseTeachers.CreateService(s.repository.getDB())
	err = dpoCourseTeachersService.UpsertMany(item.DpoCoursesTeachers)
	if err != nil {
		return err
	}
	err = dpoCourseTeachersService.DeleteMany(item.DpoCoursesTeachersForDelete)
	if err != nil {
		return err
	}
	dpoCourseDatesService := dpoCourseDates.CreateService(s.repository.getDB())
	err = dpoCourseDatesService.UpsertMany(item.DpoCoursesDates)
	if err != nil {
		return err
	}
	err = dpoCourseDatesService.DeleteMany(item.DpoCoursesDatesForDelete)
	if err != nil {
		return err
	}
	dpoCourseSpecializationsService := dpoCourseSpecializations.CreateService(s.repository.getDB())
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

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
