package children

import (
	"fmt"
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/models"
)

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

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	fmt.Println(idPool)
	return s.repository.deleteMany(idPool)
}
