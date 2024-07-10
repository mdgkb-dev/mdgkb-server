package residencyapplications

import (
	"context"
	"mdgkb/mdgkb-server/handlers/diplomas"
	"mdgkb/mdgkb-server/handlers/formvalues"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/handlers/residencyapplicationspointsachievements"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll(c context.Context) (models.ResidencyApplicationsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id *string) (*models.ResidencyApplication, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) EmailExists(c context.Context, email string, courseID string) (bool, error) {
	item, err := R.EmailExists(c, email, courseID)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) TypeExists(c context.Context, email string, main bool) (bool, error) {
	item, err := R.TypeExists(c, email, main)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) Create(c context.Context, item *models.ResidencyApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	err = diplomas.CreateService(s.helper).Upsert(item.Diploma)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	residencyApplicationsPointsAchievementsService := residencyapplicationspointsachievements.CreateService(s.helper)
	err = residencyApplicationsPointsAchievementsService.CreateMany(item.ResidencyApplicationPointsAchievements)
	if err != nil {
		return err
	}
	err = residencyApplicationsPointsAchievementsService.DeleteMany(item.ResidencyApplicationPointsAchievementsForDelete)
	if err != nil {
		return err
	}
	err = meta.S.SendApplicationsCounts()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.ResidencyApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	err = diplomas.CreateService(s.helper).Upsert(item.Diploma)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	residencyApplicationsPointsAchievementsService := residencyapplicationspointsachievements.CreateService(s.helper)
	err = residencyApplicationsPointsAchievementsService.UpsertMany(item.ResidencyApplicationPointsAchievements)
	if err != nil {
		return err
	}
	err = residencyApplicationsPointsAchievementsService.DeleteMany(item.ResidencyApplicationPointsAchievementsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateWithForm(c context.Context, item *models.FormValue) error {
	err := formvalues.CreateService(s.helper).Upsert(item)
	if err != nil {
		return err
	}
	err = R.Update(c, item.ResidencyApplication)
	if err != nil {
		return err
	}
	item.ResidencyApplication.SetIDForChildren()
	residencyApplicationsPointsAchievementsService := residencyapplicationspointsachievements.CreateService(s.helper)
	err = residencyApplicationsPointsAchievementsService.UpsertMany(item.ResidencyApplication.ResidencyApplicationPointsAchievements)
	if err != nil {
		return err
	}
	err = residencyApplicationsPointsAchievementsService.DeleteMany(item.ResidencyApplication.ResidencyApplicationPointsAchievementsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(c context.Context, items models.ResidencyApplications) error {
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
