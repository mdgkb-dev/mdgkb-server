package dailymenus

import (
	"mdgkb/mdgkb-server/handlers/dailymenuitems"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.DailyMenu) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.DailyMenu) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = dailymenuitems.CreateService(s.helper).UpsertMany(item.DailyMenuItems)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll() (models.DailyMenus, error) {
	return s.repository.getAll()
}

func (s *Service) Get(slug string) (*models.DailyMenu, error) {
	item, err := s.repository.get(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
