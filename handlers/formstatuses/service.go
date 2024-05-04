package formstatuses

import (
	"context"
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/formstatustoformstatuses"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll(c context.Context) (models.FormStatuses, error) {
	return R.GetAll(c)
}

func (s *Service) GetAllByGroupID(c context.Context, id *string) (models.FormStatuses, error) {
	return R.GetAllByGroupID(c, id)
}

func (s *Service) Get(c context.Context, id *string) (*models.FormStatus, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Upsert(c context.Context, item *models.FormStatus) error {
	err := fileinfos.CreateService(s.helper).Upsert(item.Icon)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Upsert(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = formstatustoformstatuses.CreateService(s.helper).UpsertMany(item.FormStatusToFormStatuses)
	if err != nil {
		return err
	}
	err = formstatustoformstatuses.CreateService(s.helper).DeleteMany(item.FormStatusToFormStatusesForDelete)
	return err
}

func (s *Service) UpsertMany(c context.Context, items models.FormStatuses) error {
	if len(items) == 0 {
		return nil
	}
	err := R.UpsertMany(c, items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	err = formstatustoformstatuses.CreateService(s.helper).UpsertMany(items.GetFormStatusToFormStatuses())
	if err != nil {
		return err
	}
	if len(items.GetFormStatusToFormStatusesForDelete()) > 0 {
		err = formstatustoformstatuses.CreateService(s.helper).DeleteMany(items.GetFormStatusToFormStatusesForDelete())
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Delete(c context.Context, id *string) error {
	return R.Delete(c, id)
}
