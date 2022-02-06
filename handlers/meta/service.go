package meta

import (
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
