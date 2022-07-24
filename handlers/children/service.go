package children

import (
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.Child) error {
	err := human.CreateService(s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	//items.SetIDForChildren()
	return nil
}

func (s *Service) CreateMany(items models.Children) error {
	if len(items) == 0 {
		return nil
	}
	err := human.CreateService(s.helper).CreateMany(items.GetHumans())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.createMany(items)
	if err != nil {
		return err
	}
	//items.SetIDForChildren()
	return nil
}

func (s *Service) UpsertMany(items models.Children) error {
	if len(items) == 0 {
		return nil
	}
	err := human.CreateService(s.helper).UpsertMany(items.GetHumans())
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
	err := human.CreateService(s.helper).Upsert(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	if item == nil {
		return nil
	}
	err = s.repository.upsert(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
