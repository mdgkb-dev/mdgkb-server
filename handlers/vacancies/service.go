package vacancies

import (
	"mdgkb/mdgkb-server/handlers/vacancyduties"
	"mdgkb/mdgkb-server/handlers/vacancyrequirements"
	"mdgkb/mdgkb-server/handlers/vacancyresponse"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Vacancy) error {
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	vacancyDutiesService := vacancyduties.CreateService(s.helper)
	err = vacancyDutiesService.UpsertMany(item.VacancyDuties)
	if err != nil {
		return err
	}
	vacancyRequirementsService := vacancyrequirements.CreateService(s.helper)
	err = vacancyRequirementsService.UpsertMany(item.VacancyRequirements)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.VacanciesWithCount, error) {
	return s.repository.getAll()
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
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	vacancyDutiesService := vacancyduties.CreateService(s.helper)
	err = vacancyDutiesService.UpsertMany(item.VacancyDuties)
	if err != nil {
		return err
	}
	err = vacancyDutiesService.DeleteMany(item.VacancyDutiesDelete)
	if err != nil {
		return err
	}

	vacancyRequirementsService := vacancyrequirements.CreateService(s.helper)
	err = vacancyRequirementsService.UpsertMany(item.VacancyRequirements)
	if err != nil {
		return err
	}
	err = vacancyRequirementsService.DeleteMany(item.VacancyDutiesDelete)
	if err != nil {
		return err
	}

	err = vacancyresponse.CreateService(s.helper).DeleteMany(item.VacancyResponsesForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) CreateResponse(item *models.VacancyResponse) error {
	//err := human.CreateService(s.helper).Create(item.Human)
	//if err != nil {
	//	return err
	//}
	item.SetForeignKeys()
	err := s.repository.createResponse(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	return nil
}

func (s *Service) setQueryFilter(c *gin.Context) error {
	return s.repository.setQueryFilter(c)
}

func (s *Service) UpdateMany(items models.Vacancies) error {
	items.SetForeignKeys()
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}
