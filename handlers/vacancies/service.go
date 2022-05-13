package vacancies

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/handlers/vacancyDuties"
	"mdgkb/mdgkb-server/handlers/vacancyRequirements"
	"mdgkb/mdgkb-server/handlers/vacancyResponsesToDocuments"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Vacancy) error {
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	vacancyDutiesService := vacancyDuties.CreateService(s.repository.GetDB())
	err = vacancyDutiesService.UpsertMany(item.VacancyDuties)
	if err != nil {
		return err
	}
	vacancyRequirementsService := vacancyRequirements.CreateService(s.repository.GetDB())
	err = vacancyRequirementsService.UpsertMany(item.VacancyRequirements)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.Vacancies, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.Vacancy, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetBySlug(slug *string) (*models.Vacancy, error) {
	return s.repository.getBySlug(slug)
}

func (s *Service) Update(item *models.Vacancy) error {
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	vacancyDutiesService := vacancyDuties.CreateService(s.repository.GetDB())
	err = vacancyDutiesService.UpsertMany(item.VacancyDuties)
	if err != nil {
		return err
	}
	err = vacancyDutiesService.DeleteMany(item.VacancyDutiesDelete)
	if err != nil {
		return err
	}
	vacancyRequirementsService := vacancyRequirements.CreateService(s.repository.GetDB())
	err = vacancyRequirementsService.UpsertMany(item.VacancyRequirements)
	if err != nil {
		return err
	}
	err = vacancyRequirementsService.DeleteMany(item.VacancyDutiesDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) CreateResponse(item *models.VacancyResponse) error {
	//err := human.CreateService(s.repository.GetDB(), s.helper).Create(item.Human)
	//if err != nil {
	//	return err
	//}
	item.SetForeignKeys()
	err := s.repository.createResponse(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = vacancyResponsesToDocuments.CreateService(s.repository.GetDB()).CreateMany(item.VacancyResponsesToDocuments)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SetQueryFilter(c *gin.Context) error {
	return s.repository.SetQueryFilter(c)
}
