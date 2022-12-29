package dishessamples

import (
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.DishSample) error {
	err := fileinfos.CreateService(s.helper).Upsert(item.Image)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.DishSample) error {
	err := fileinfos.CreateService(s.helper).Upsert(item.Image)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.DishSamples, error) {
	return s.repository.getAll()
}

func (s *Service) Get(slug string) (*models.DishSample, error) {
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

func (s *Service) UpdateAll(items models.DishSamples) error {
	err := s.repository.updateAll(items)
	if err != nil {
		return err
	}
	return nil
}
