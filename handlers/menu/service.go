package menu

import (
	"fmt"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/subMenus"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Menu) error {
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(models.FileInfos{item.Icon})
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = subMenus.CreateService(s.repository.getDB()).CreateMany(item.SubMenus)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.Menus, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.Menu, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Menu) error {
	fmt.Println(item.Icon)
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(models.FileInfos{item.Icon})
	if err != nil {
		return err
	}
	item.SetForeignKeys()

	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	subMenuService := subMenus.CreateService(s.repository.getDB())
	err = subMenuService.DeleteMany(item.SubMenusForDelete)
	if err != nil {
		return err
	}
	err = subMenuService.UpsertMany(item.SubMenus)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
