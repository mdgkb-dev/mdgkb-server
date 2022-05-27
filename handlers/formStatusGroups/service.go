package formStatusGroups

import (
	"mdgkb/mdgkb-server/models"
	"github.com/gin-gonic/gin"
)

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) GetAll() (models.FormStatusGroupsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.FormStatusGroup, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Upsert(item *models.FormStatusGroup) error {
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.FormStatusGroups) error {
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
