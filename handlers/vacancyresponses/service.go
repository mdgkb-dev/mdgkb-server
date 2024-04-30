package vacancyresponses

import (
	"context"
	"mdgkb/mdgkb-server/handlers/formvalues"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(c context.Context, item *models.VacancyResponse) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.VacancyResponsesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.VacancyResponse, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.VacancyResponse) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Update(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) DeleteMany(c context.Context, idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return R.DeleteMany(c, idPool)
}

func (s *Service) EmailExists(c context.Context, email string, vacancyID string) (bool, error) {
	item, err := R.EmailExists(c, email, vacancyID)
	if err != nil {
		return item, err
	}
	return item, nil
}
