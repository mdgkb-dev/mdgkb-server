package forms

import (
	"mdgkb/mdgkb-server/handlers/fields"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Form) error {
	if item == nil {
		return nil
	}
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = fields.CreateService(s.helper).UpsertMany(item.Fields)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Form) error {
	if item == nil {
		return nil
	}
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = fields.CreateService(s.helper).UpsertMany(item.Fields)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.Forms) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Upsert(item *models.Form) error {
	if item == nil {
		return nil
	}
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = fields.CreateService(s.helper).UpsertMany(item.Fields)
	if err != nil {
		return err
	}
	return nil
}

//func (s *Service) DeleteMany(idPool []string) error {
//	if len(idPool) == 0 {
//		return nil
//	}
//	return s.repository.deleteMany(idPool)
//}
