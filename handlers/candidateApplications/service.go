package candidateApplications

import (
	"mdgkb/mdgkb-server/handlers/candidateApplicationSpecializations"
	"mdgkb/mdgkb-server/handlers/fieldsValues"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.CandidateApplications, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.CandidateApplication, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.CandidateApplication) error {
	err := users.CreateService(s.repository.getDB(), s.helper).UpsertEmail(item.User)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = fieldsValues.CreateService(s.repository.getDB()).UpsertMany(item.FieldValues)
	if err != nil {
		return err
	}
	err = candidateApplicationSpecializations.CreateService(s.repository.getDB()).UpsertMany(item.CandidateApplicationSpecializations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.CandidateApplication) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = fieldsValues.CreateService(s.repository.getDB()).UpsertMany(item.FieldValues)
	if err != nil {
		return err
	}
	candidateApplicationSpecializationsService := candidateApplicationSpecializations.CreateService(s.repository.getDB())
	err = candidateApplicationSpecializationsService.UpsertMany(item.CandidateApplicationSpecializations)
	if err != nil {
		return err
	}
	err = candidateApplicationSpecializationsService.DeleteMany(item.CandidateApplicationSpecializationsForDelete)
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
