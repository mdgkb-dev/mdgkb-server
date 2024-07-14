package educationalacademics

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) GetAll(c context.Context) (models.EducationalAcademics, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.EducationalAcademic, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(c context.Context, item *models.EducationalAcademic) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.EducationalAcademic) error {
	if item == nil {
		return nil
	}
	err := R.Upsert(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) UpdateAll(c context.Context, items models.EducationalAcademics) error {
	return R.UpdateAll(c, items)
}

func (s *Service) DeleteMany(c context.Context, idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return R.DeleteMany(c, idPool)
}
