package visitsapplications

import (
	"context"
	"mdgkb/mdgkb-server/handlers/formvalues"
	"mdgkb/mdgkb-server/handlers/visits"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll(c context.Context) (models.VisitsApplicationsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id *string) (*models.VisitsApplication, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(c context.Context, item *models.VisitsApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = visits.CreateService(s.helper).UpsertMany(item.Visits)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.VisitsApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Update(c, item)
	if err != nil {
		return err
	}
	err = visits.CreateService(s.helper).UpsertMany(item.Visits)
	if err != nil {
		return err
	}
	if len(item.VisitsForDelete) > 0 {
		err = visits.CreateService(s.helper).DeleteMany(item.VisitsForDelete)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Delete(c context.Context, id *string) error {
	return R.Delete(c, id)
}
