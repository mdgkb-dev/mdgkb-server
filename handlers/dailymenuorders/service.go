package dailymenuorders

import (
	"mdgkb/mdgkb-server/handlers/dailymenuorderitems"
	"mdgkb/mdgkb-server/handlers/formvalues"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.DailyMenuOrder) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	dailyMenuOrderItemsService := dailymenuorderitems.CreateService(s.helper)
	err = dailyMenuOrderItemsService.UpsertMany(item.DailyMenuOrderItems)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.DailyMenuOrder) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	dailyMenuOrderItemsService := dailymenuorderitems.CreateService(s.helper)
	err = dailyMenuOrderItemsService.UpsertMany(item.DailyMenuOrderItems)
	if err != nil {
		return err
	}
	err = dailyMenuOrderItemsService.DeleteMany(item.DailyMenuOrderItemsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.DailyMenuOrdersWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(slug string) (*models.DailyMenuOrder, error) {
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

func (s *Service) UpsertMany(items models.DailyMenuOrders) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}
