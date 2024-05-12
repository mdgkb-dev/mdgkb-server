package dailymenus

import (
	"context"
	"mdgkb/mdgkb-server/handlers/dailymenuitems"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(c context.Context, item *models.DailyMenu) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.DailyMenu) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = dailymenuitems.CreateService(s.helper).UpsertMany(item.DailyMenuItems)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateAll(c context.Context, items models.DailyMenus) error {
	err := R.UpdateAll(c, items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.DailyMenus, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, slug string) (*models.DailyMenu, error) {
	return R.Get(c, slug)
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) GetTodayActive(c context.Context) (*models.DailyMenu, error) {
	return R.GetTodayActive(c)
}
