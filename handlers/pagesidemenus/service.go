package pagesidemenus

import (
	"mdgkb/mdgkb-server/handlers/pagesections"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.PageSideMenus, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.PageSideMenu, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.PageSideMenu) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = pagesections.CreateService(s.helper).UpsertMany(item.PageSections)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Update(item *models.PageSideMenu) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	documentTypeService := pagesections.CreateService(s.helper)
	err = documentTypeService.DeleteMany(item.PageSectionsForDelete)
	if err != nil {
		return err
	}
	err = documentTypeService.UpsertMany(item.PageSections)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) UpdateOrder(items models.PageSideMenus) error {
	return s.repository.upsertMany(items)
}
