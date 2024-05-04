package formstatusgroups

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll(c context.Context) (models.FormStatusGroupsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id *string) (*models.FormStatusGroup, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Upsert(c context.Context, item *models.FormStatusGroup) error {
	err := R.Upsert(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(c context.Context, items models.FormStatusGroups) error {
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
