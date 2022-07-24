package residencycourses

import (
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/residencycoursespecializations"
	"mdgkb/mdgkb-server/handlers/residencycourseteachers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.ResidencyCoursesWithCount, error) {
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
	fileInfosService := fileinfos.CreateService(s.helper)
	err := fileInfosService.Create(item.Program)
	if err != nil {
		return err
	}
	err = fileInfosService.Create(item.Annotation)
	if err != nil {
		return err
	}
	err = fileInfosService.Create(item.Schedule)
	if err != nil {
		return err
	}
	err = fileInfosService.Create(item.Plan)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = residencycourseteachers.CreateService(s.helper).UpsertMany(item.ResidencyCoursesTeachers)
	if err != nil {
		return err
	}
	err = residencycoursespecializations.CreateService(s.helper).UpsertMany(item.ResidencyCoursesSpecializations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.ResidencyCourse) error {
	fileInfosService := fileinfos.CreateService(s.helper)
	err := fileInfosService.Upsert(item.Program)
	if err != nil {
		return err
	}
	err = fileInfosService.Upsert(item.Annotation)
	if err != nil {
		return err
	}
	err = fileInfosService.Upsert(item.Schedule)
	if err != nil {
		return err
	}
	err = fileInfosService.Upsert(item.Plan)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	residencyCourseTeachersService := residencycourseteachers.CreateService(s.helper)
	err = residencyCourseTeachersService.UpsertMany(item.ResidencyCoursesTeachers)
	if err != nil {
		return err
	}
	err = residencyCourseTeachersService.DeleteMany(item.ResidencyCoursesTeachersForDelete)
	if err != nil {
		return err
	}
	residencyCourseSpecializationsService := residencycoursespecializations.CreateService(s.helper)
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

func (s *Service) UpsertMany(items models.ResidencyCourses) error {
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
