package vacancies

import (
	"context"
	"mdgkb/mdgkb-server/handlers/vacancyduties"
	"mdgkb/mdgkb-server/handlers/vacancyrequirements"
	"mdgkb/mdgkb-server/handlers/vacancyresponses"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(c context.Context, item *models.Vacancy) error {
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	err := R.Create(c, item)
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

func (s *Service) GetAll(c context.Context) (models.VacanciesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id *string) (*models.Vacancy, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetBySlug(c context.Context, slug *string) (*models.Vacancy, error) {
	return R.GetBySlug(c, slug)
}

func (s *Service) Update(c context.Context, item *models.Vacancy) error {
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	err := R.Update(c, item)
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

	err = vacancyresponses.S.DeleteMany(c, item.VacancyResponsesForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(c context.Context, id *string) error {
	return R.Delete(c, id)
}

func (s *Service) CreateResponse(c context.Context, item *models.VacancyResponse) error {
	//err := human.CreateService(s.helper).Create(item.Human)
	//if err != nil {
	//	return err
	//}
	item.SetForeignKeys()
	err := R.CreateResponse(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	return nil
}

func (s *Service) UpdateMany(c context.Context, items models.Vacancies) error {
	items.SetForeignKeys()
	err := R.UpsertMany(c, items)
	if err != nil {
		return err
	}
	return nil
}
