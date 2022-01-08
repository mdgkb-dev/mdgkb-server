package meta

import "mdgkb/mdgkb-server/models/schema"

func (s *Service) GetCount(table *string) (*int, error) {
	return s.repository.getCount(table)
}

func (s *Service) GetSchema() schema.Schema {
	return schema.CreateSchema()
}
