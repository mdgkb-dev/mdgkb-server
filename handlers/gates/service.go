package gates

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.Gates, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.Gate, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.Gate) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Gate) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateMany(items models.Gates) error {
	items.SetForeignKeys()
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
