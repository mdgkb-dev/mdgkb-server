package postgraduatecourses

import (
	"context"
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/postgraduatecoursedates"
	"mdgkb/mdgkb-server/handlers/postgraduatecourseplans"
	"mdgkb/mdgkb-server/handlers/postgraduatecoursespecializations"
	"mdgkb/mdgkb-server/handlers/postgraduatecourseteachers"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll(c context.Context) (models.PostgraduateCoursesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context) (*models.PostgraduateCourse, error) {
	item, err := R.Get(c)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(c context.Context, item *models.PostgraduateCourse) error {
	fileInfosService := fileinfos.CreateService(s.helper)
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
	err = fileInfosService.Create(item.Annotation)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = postgraduatecourseteachers.CreateService(s.helper).UpsertMany(item.PostgraduateCoursesTeachers)
	if err != nil {
		return err
	}
	err = postgraduatecoursedates.CreateService(s.helper).UpsertMany(item.PostgraduateCoursesDates)
	if err != nil {
		return err
	}
	err = postgraduatecoursespecializations.CreateService(s.helper).UpsertMany(item.PostgraduateCoursesSpecializations)
	if err != nil {
		return err
	}
	postgraduateCoursePlansService := postgraduatecourseplans.CreateService(s.helper)
	err = postgraduateCoursePlansService.UpsertMany(item.PostgraduateCoursePlans)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.PostgraduateCourse) error {
	fileInfosService := fileinfos.CreateService(s.helper)
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
	err = fileInfosService.Upsert(item.Annotation)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	postgraduateCourseTeachersService := postgraduatecourseteachers.CreateService(s.helper)
	err = postgraduateCourseTeachersService.UpsertMany(item.PostgraduateCoursesTeachers)
	if err != nil {
		return err
	}
	err = postgraduateCourseTeachersService.DeleteMany(item.PostgraduateCoursesTeachersForDelete)
	if err != nil {
		return err
	}
	postgraduateCourseDatesService := postgraduatecoursedates.CreateService(s.helper)
	err = postgraduateCourseDatesService.UpsertMany(item.PostgraduateCoursesDates)
	if err != nil {
		return err
	}
	err = postgraduateCourseDatesService.DeleteMany(item.PostgraduateCoursesDatesForDelete)
	if err != nil {
		return err
	}
	postgraduateCourseSpecializationsService := postgraduatecoursespecializations.CreateService(s.helper)
	err = postgraduateCourseSpecializationsService.UpsertMany(item.PostgraduateCoursesSpecializations)
	if err != nil {
		return err
	}
	err = postgraduateCourseSpecializationsService.DeleteMany(item.PostgraduateCoursesSpecializationsForDelete)
	if err != nil {
		return err
	}
	postgraduateCoursePlansService := postgraduatecourseplans.CreateService(s.helper)
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

func (s *Service) UpsertMany(c context.Context, items models.PostgraduateCourses) error {
	if len(items) == 0 {
		return nil
	}
	err := R.UpsertMany(c, items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(c context.Context, id *string) error {
	return R.Delete(c, id)
}
