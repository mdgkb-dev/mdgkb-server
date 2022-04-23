package meta

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/models/schema"
)

func (s *Service) GetCount(table *string) (*int, error) {
	return s.repository.getCount(table)
}

func (s *Service) GetSchema() schema.Schema {
	return schema.CreateSchema()
}

func (s *Service) GetOptions(optionModel *models.OptionModel) (models.Options, error) {
	return s.repository.getOptions(optionModel)
}

func (s *Service) SendApplicationsCounts() error {
	items, err := s.repository.getApplicationsCounts()
	if err != nil {
		return err
	}
	fmt.Println(items)
	fmt.Println(items)
	fmt.Println(items)
	fmt.Println(items)
	s.helper.Broker.SendEvent("applications-counts-get", items)
	return nil
}

func (s *Service) GetApplicationsCounts() (models.ApplicationsCounts, error) {
	return s.repository.getApplicationsCounts()
}
