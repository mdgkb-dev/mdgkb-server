package dpoApplications

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.DpoApplications, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.DpoApplication, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.DpoApplication) error {
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(item.GetFileInfos())
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

func (s *Service) Update(item *models.DpoApplication) error {
	err := s.repository.update(item)
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
