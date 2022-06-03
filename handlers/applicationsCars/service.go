package applicationsCars

import (
	"mdgkb/mdgkb-server/handlers/formValues"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.ApplicationsCarsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.ApplicationCar, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.ApplicationCar) error {
	err := formValues.CreateService(s.repository.GetDB(), s.helper).Upsert(item.FormValue)
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

func (s *Service) Update(item *models.ApplicationCar) error {
	err := formValues.CreateService(s.repository.GetDB(), s.helper).Upsert(item.FormValue)
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

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) error {
	return s.repository.SetQueryFilter(c)
}
