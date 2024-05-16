package educationalmanagers

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll(c context.Context) (models.EducationalManagers, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.EducationalManager, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(c context.Context, item *models.EducationalManager) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.EducationalManager) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
