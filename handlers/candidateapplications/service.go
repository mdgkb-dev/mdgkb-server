package candidateapplications

import (
	"mdgkb/mdgkb-server/handlers/candidateapplicationspecializations"
	"mdgkb/mdgkb-server/handlers/formvalues"
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

func (s *Service) EmailExists(email string, examID string) (bool, error) {
	item, err := s.repository.emailExists(email, examID)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) Create(item *models.CandidateApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = candidateapplicationspecializations.CreateService(s.helper).UpsertMany(item.CandidateApplicationSpecializations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.CandidateApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	candidateApplicationSpecializationsService := candidateapplicationspecializations.CreateService(s.helper)
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
