package meta

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetCount(table *string) (*int, error) {
	return s.repository.getCount(table)
}

func (s *Service) GetOptions(optionModel *models.OptionModel) (models.Options, error) {
	return s.repository.getOptions(optionModel)
}

func (s *Service) SendApplicationsCounts() error {
	items, err := s.repository.getApplicationsCounts()
	if err != nil {
		return err
	}
	s.helper.Broker.SendEvent("applications-counts-get", items)
	return nil
}

func (s *Service) GetApplicationsCounts() (models.ApplicationsCounts, error) {
	return s.repository.getApplicationsCounts()
}
