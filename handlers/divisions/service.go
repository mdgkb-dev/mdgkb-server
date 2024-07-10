package divisions

import (
	"context"
	"mdgkb/mdgkb-server/handlers/divisionimages"
	"mdgkb/mdgkb-server/handlers/divisionvideos"
	"mdgkb/mdgkb-server/handlers/doctorsdivisions"
	"mdgkb/mdgkb-server/handlers/schedules"
	"mdgkb/mdgkb-server/handlers/visitingrulesgroups"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(c context.Context, item *models.Division) error {
	// timetableService := timetables.CreateService(s.helper)
	// err := timetableService.Create(item.Timetable)
	// if err != nil {
	// 	return err
	// }
	schedulesService := schedules.CreateService(s.helper)
	err := schedulesService.Create(item.Schedule)
	if err != nil {
		return err
	}
	// item.Slug = s.helper.Util.MakeSlug(item.Name, true)

	// contactInfoService := contactinfo.CreateService(s.helper)
	// err = contactInfoService.Create(item.ContactInfo)
	// if err != nil {
	// 	return err
	// }
	item.SetForeignKeys()

	doctorsDivisionsService := doctorsdivisions.CreateService(s.helper)
	err = doctorsDivisionsService.UpsertMany(item.DoctorsDivisions)
	if err != nil {
		return err
	}

	err = R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	visitingRulesGroupsService := visitingrulesgroups.CreateService(s.helper)
	err = visitingRulesGroupsService.UpsertMany(item.VisitingRulesGroups)
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

func (s *Service) Update(c context.Context, item *models.Division) error {
	// timetableService := timetables.CreateService(s.helper)
	// err := timetableService.Upsert(item.Timetable)
	// if err != nil {
	// 	return err
	// }

	schedulesService := schedules.CreateService(s.helper)
	err := schedulesService.Upsert(item.Schedule)
	if err != nil {
		return err
	}

	// contactInfoService := contactinfo.CreateService(s.helper)
	// err = contactInfoService.Upsert(item.ContactInfo)
	// if err != nil {
	// 	return err
	// }
	// item.SetForeignKeys()

	item.SetIDForChildren()
	doctorsDivisionsService := doctorsdivisions.CreateService(s.helper)
	err = doctorsDivisionsService.DeleteMany(item.DoctorsDivisionsForDelete)
	if err != nil {
		return err
	}
	err = doctorsDivisionsService.UpsertMany(item.DoctorsDivisions)
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
	visitingRulesGroupsService := visitingrulesgroups.CreateService(s.helper)
	err = visitingRulesGroupsService.UpsertMany(item.VisitingRulesGroups)
	if err != nil {
		return err
	}
	err = visitingRulesGroupsService.DeleteMany(item.VisitingRulesGroupsForDelete)
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
	return R.Update(c, item)
}

func (s *Service) GetAll(c context.Context) (models.DivisionsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Division, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) GetBySearch(c context.Context, search string) (models.Divisions, error) {
	return R.GetBySearch(c, search)
}
