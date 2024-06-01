package residencycourses

import (
	"context"
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/residencycoursepracticeplacegroups"
	"mdgkb/mdgkb-server/handlers/residencycoursespecializations"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll(c context.Context) (models.ResidencyCoursesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.ResidencyCourse, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(c context.Context, item *models.ResidencyCourse) error {
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
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = residencycoursespecializations.CreateService(s.helper).UpsertMany(item.ResidencyCoursesSpecializations)
	if err != nil {
		return err
	}
	err = residencycoursepracticeplacegroups.CreateService(s.helper).CreateMany(item.ResidencyCoursePracticePlaceGroups)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.ResidencyCourse) error {
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
	err = R.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	residencyCourseSpecializationsService := residencycoursespecializations.CreateService(s.helper)
	err = residencyCourseSpecializationsService.UpsertMany(item.ResidencyCoursesSpecializations)
	if err != nil {
		return err
	}
	err = residencyCourseSpecializationsService.DeleteMany(item.ResidencyCoursesSpecializationsForDelete)
	if err != nil {
		return err
	}
	residencyCoursePracticePlaceGroupsService := residencycoursepracticeplacegroups.CreateService(s.helper)
	err = residencyCoursePracticePlaceGroupsService.UpsertMany(item.ResidencyCoursePracticePlaceGroups)
	if err != nil {
		return err
	}
	err = residencyCoursePracticePlaceGroupsService.DeleteMany(item.ResidencyCoursePracticePlaceGroupsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(c context.Context, items models.ResidencyCourses) error {
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
