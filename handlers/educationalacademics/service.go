package educationalacademics

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) GetAll() (models.EducationalAcademics, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.EducationalAcademic, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.EducationalAcademic) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.EducationalAcademic) error {
	if item == nil {
		return nil
	}
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) UpdateAll(items models.EducationalAcademics) error {
	return s.repository.updateAll(items)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
