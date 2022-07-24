package menus

import (
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/submenus"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Menu) error {
	err := fileinfos.CreateService(s.helper).UpsertMany(models.FileInfos{item.Icon})
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = submenus.CreateService(s.helper).CreateMany(item.SubMenus)
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
	err := fileinfos.CreateService(s.helper).UpsertMany(models.FileInfos{item.Icon})
	if err != nil {
		return err
	}
	item.SetForeignKeys()

	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	subMenuService := submenus.CreateService(s.helper)
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

func (s *Service) UpsertMany(items WithDeleted) error {
	err := fileinfos.CreateService(s.helper).UpsertMany(items.Menus.GetIcons())
	if err != nil {
		return err
	}
	items.Menus.SetIDForChildren()

	err = s.repository.upsertMany(items.Menus)
	if err != nil {
		return err
	}
	if len(items.MenusForDeleted) > 0 {
		err = s.repository.deleteMany(items.MenusForDeleted)
		if err != nil {
			return err
		}
		items.Menus.SetIDForChildren()
	}
	subMenuService := submenus.CreateService(s.helper)
	err = subMenuService.DeleteMany(items.Menus.GetSubMenusForDelete())
	if err != nil {
		return err
	}
	err = subMenuService.UpsertMany(items.Menus.GetSubMenus())
	if err != nil {
		return err
	}
	return nil
}
