package faqs

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(c context.Context, item *models.Faq) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.Faqs, error) {
	items, err := R.GetAll(c)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(c context.Context, id string) (*models.Faq, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) UpsertMany(c context.Context, items models.Faqs) error {
	err := R.UpsertMany(c, items)
	return err
}

func (s *Service) Update(c context.Context, item *models.Faq) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
