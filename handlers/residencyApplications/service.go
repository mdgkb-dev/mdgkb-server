package residencyApplications

import (
	"mdgkb/mdgkb-server/handlers/formValues"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.ResidencyApplicationsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.ResidencyApplication, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) EmailExists(email string, courseId string) (bool, error) {
	item, err := s.repository.emailExists(email, courseId)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) Create(item *models.ResidencyApplication) error {
	err := formValues.CreateService(s.repository.getDB(), s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	err = meta.CreateService(s.repository.getDB(), s.helper).SendApplicationsCounts()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.ResidencyApplication) error {
	err := formValues.CreateService(s.repository.getDB(), s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	return nil
}

func (s *Service) UpsertMany(items models.ResidencyApplications) error {
	if len(items) == 0 {
		return nil
	}
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
