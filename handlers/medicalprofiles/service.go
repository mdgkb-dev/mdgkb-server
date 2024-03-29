package medicalprofiles

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.MedicalProfile) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.MedicalProfile) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.MedicalProfiles, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.MedicalProfile, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
