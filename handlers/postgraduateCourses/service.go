package postgraduateCourses

import (
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/postgraduateCourseDates"
	"mdgkb/mdgkb-server/handlers/postgraduateCoursePlans"
	"mdgkb/mdgkb-server/handlers/postgraduateCourseSpecializations"
	"mdgkb/mdgkb-server/handlers/postgraduateCourseTeachers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.PostgraduateCourses, error) {
	return s.repository.getAll()
}

func (s *Service) Get() (*models.PostgraduateCourse, error) {
	item, err := s.repository.get()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.PostgraduateCourse) error {
	fileInfosService := fileInfos.CreateService(s.repository.getDB())
	err := fileInfosService.Create(item.QuestionsFile)
	if err != nil {
		return err
	}
	err = fileInfosService.Create(item.ProgramFile)
	if err != nil {
		return err
	}
	err = fileInfosService.Create(item.Calendar)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = postgraduateCourseTeachers.CreateService(s.repository.getDB()).UpsertMany(item.PostgraduateCoursesTeachers)
	if err != nil {
		return err
	}
	err = postgraduateCourseDates.CreateService(s.repository.getDB()).UpsertMany(item.PostgraduateCoursesDates)
	if err != nil {
		return err
	}
	err = postgraduateCourseSpecializations.CreateService(s.repository.getDB()).UpsertMany(item.PostgraduateCoursesSpecializations)
	if err != nil {
		return err
	}
	postgraduateCoursePlansService := postgraduateCoursePlans.CreateService(s.repository.getDB())
	err = postgraduateCoursePlansService.UpsertMany(item.PostgraduateCoursePlans)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.PostgraduateCourse) error {
	fileInfosService := fileInfos.CreateService(s.repository.getDB())
	err := fileInfosService.Upsert(item.QuestionsFile)
	if err != nil {
		return err
	}
	err = fileInfosService.Upsert(item.ProgramFile)
	if err != nil {
		return err
	}
	err = fileInfosService.Upsert(item.Calendar)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	postgraduateCourseTeachersService := postgraduateCourseTeachers.CreateService(s.repository.getDB())
	err = postgraduateCourseTeachersService.UpsertMany(item.PostgraduateCoursesTeachers)
	if err != nil {
		return err
	}
	err = postgraduateCourseTeachersService.DeleteMany(item.PostgraduateCoursesTeachersForDelete)
	if err != nil {
		return err
	}
	postgraduateCourseDatesService := postgraduateCourseDates.CreateService(s.repository.getDB())
	err = postgraduateCourseDatesService.UpsertMany(item.PostgraduateCoursesDates)
	if err != nil {
		return err
	}
	err = postgraduateCourseDatesService.DeleteMany(item.PostgraduateCoursesDatesForDelete)
	if err != nil {
		return err
	}
	postgraduateCourseSpecializationsService := postgraduateCourseSpecializations.CreateService(s.repository.getDB())
	err = postgraduateCourseSpecializationsService.UpsertMany(item.PostgraduateCoursesSpecializations)
	if err != nil {
		return err
	}
	err = postgraduateCourseSpecializationsService.DeleteMany(item.PostgraduateCoursesSpecializationsForDelete)
	if err != nil {
		return err
	}
	postgraduateCoursePlansService := postgraduateCoursePlans.CreateService(s.repository.getDB())
	err = postgraduateCoursePlansService.UpsertMany(item.PostgraduateCoursePlans)
	if err != nil {
		return err
	}
	err = postgraduateCoursePlansService.DeleteMany(item.PostgraduateCoursePlansForDelete)
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
