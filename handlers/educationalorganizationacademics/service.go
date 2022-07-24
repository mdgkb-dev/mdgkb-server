package educationalorganizationacademics

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) GetAll() (models.EducationalOrganizationAcademics, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpsertMany(items models.EducationalOrganizationAcademics) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}

func (s *Service) Upsert(item *models.EducationalOrganizationAcademic) error {
	if item == nil {
		return nil
	}
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(id []string) error {
	if len(id) == 0 {
		return nil
	}
	return s.repository.deleteMany(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) DeleteByDoctorID(id uuid.NullUUID) error {
	if !id.Valid {
		return nil
	}
	return s.repository.deleteByDoctorID(id)
}
