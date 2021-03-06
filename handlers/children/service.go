package children

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Child) error {
	err := human.CreateService(s.repository.getDB(), s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	//items.SetIdForChildren()
	return nil
}

func (s *Service) CreateMany(items models.Children) error {
	if len(items) == 0 {
		return nil
	}
	err := human.CreateService(s.repository.getDB(), s.helper).CreateMany(items.GetHumans())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.createMany(items)
	if err != nil {
		return err
	}
	//items.SetIdForChildren()
	return nil
}

func (s *Service) UpsertMany(items models.Children) error {
	if len(items) == 0 {
		return nil
	}
	err := human.CreateService(s.repository.getDB(), s.helper).UpsertMany(items.GetHumans())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Upsert(item *models.Child) error {
	if item == nil {
		return nil
	}
	return s.repository.upsert(item)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
