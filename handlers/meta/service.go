package meta

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetCount(table *string) (*int, error) {
	return R.getCount(table)
}

func (s *Service) GetOptions(optionModel *models.OptionModel) (models.Options, error) {
	return R.getOptions(optionModel)
}

func (s *Service) SendApplicationsCounts() error {
	items, err := R.getApplicationsCounts()
	if err != nil {
		return err
	}
	s.helper.Broker.SendEvent("applications-counts-get", items)
	return nil
}

func (s *Service) GetApplicationsCounts() (models.ApplicationsCounts, error) {
	return R.getApplicationsCounts()
}

func (s *Service) SearchMain(c context.Context, searchModel *models.SearchModel) (err error) {
	searchModel.SearchGroups, err = R.GetGroups(c, searchModel.SearchGroupID)
	if err != nil {
		return err
	}
	for i := range searchModel.SearchGroups {
		err = R.Search(c, searchModel.SearchGroups[i], searchModel)
		if err != nil {
			return err
		}
		searchModel.SearchGroups[i].BuildRoutes()
	}
	return nil
}
