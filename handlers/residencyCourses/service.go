package residencyCourses

import (
	"mdgkb/mdgkb-server/handlers/residencyCourseSpecializations"
	"mdgkb/mdgkb-server/handlers/residencyCourseTeachers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.ResidencyCourses, error) {
	return s.repository.getAll()
}

func (s *Service) Get() (*models.ResidencyCourse, error) {
	item, err := s.repository.get()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.ResidencyCourse) error {
	item.SetForeignKeys()
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = residencyCourseTeachers.CreateService(s.repository.getDB()).UpsertMany(item.ResidencyCoursesTeachers)
	if err != nil {
		return err
	}
	err = residencyCourseSpecializations.CreateService(s.repository.getDB()).UpsertMany(item.ResidencyCoursesSpecializations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.ResidencyCourse) error {
	item.SetForeignKeys()
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	residencyCourseTeachersService := residencyCourseTeachers.CreateService(s.repository.getDB())
	err = residencyCourseTeachersService.UpsertMany(item.ResidencyCoursesTeachers)
	if err != nil {
		return err
	}
	err = residencyCourseTeachersService.DeleteMany(item.ResidencyCoursesTeachersForDelete)
	if err != nil {
		return err
	}
	residencyCourseSpecializationsService := residencyCourseSpecializations.CreateService(s.repository.getDB())
	err = residencyCourseSpecializationsService.UpsertMany(item.ResidencyCoursesSpecializations)
	if err != nil {
		return err
	}
	err = residencyCourseSpecializationsService.DeleteMany(item.ResidencyCoursesSpecializationsForDelete)
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
