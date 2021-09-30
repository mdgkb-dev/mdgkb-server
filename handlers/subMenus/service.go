package subMenus

import (
	"mdgkb/mdgkb-server/handlers/subSubMenus"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.SubMenus) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	subSubMenuService := subSubMenus.CreateService(s.repository.getDB())
	err = subSubMenuService.CreateMany(items.GetSubSubMenus())
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.SubMenus) error {

	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	subSubMenuService := subSubMenus.CreateService(s.repository.getDB())
	err = subSubMenuService.DeleteMany(items.GetIDForDelete())
	if err != nil{
		return err
	}
	err = subSubMenuService.UpsertMany(items.GetSubSubMenus())
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}

