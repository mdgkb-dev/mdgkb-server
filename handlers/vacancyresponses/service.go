package vacancyresponses

import (
	"mdgkb/mdgkb-server/handlers/formvalues"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) Create(item *models.VacancyResponse) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
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

func (s *Service) GetAll() (models.VacancyResponsesWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.VacancyResponse, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.VacancyResponse) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
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

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}

func (s *Service) EmailExists(email string, vacancyID string) (bool, error) {
	item, err := s.repository.emailExists(email, vacancyID)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) setQueryFilter(c *gin.Context) error {
	return s.repository.setQueryFilter(c)
}
