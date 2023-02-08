package pagesidemenus

import (
	"mdgkb/mdgkb-server/handlers/pagesections"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (s *Service) UpsertMany(items models.PageSideMenus) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	pageSectionsService := pagesections.CreateService(s.helper)
	err = pageSectionsService.DeleteMany(items.GetPageSectionsForDelete())
	if err != nil {
		return err
	}
	err = pageSectionsService.UpsertMany(items.GetPageSections())
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
