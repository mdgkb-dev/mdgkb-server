package formpatterns

import (
	"context"
	"mdgkb/mdgkb-server/handlers/fields"
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll(c context.Context) (models.FormPatterns, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.FormPattern, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(c context.Context, item *models.FormPattern) error {
	err := fileinfos.CreateService(s.helper).Create(item.PersonalDataAgreement)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = fields.CreateService(s.helper).UpsertMany(item.Fields)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.FormPattern) error {
	err := fileinfos.CreateService(s.helper).Upsert(item.PersonalDataAgreement)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	fieldsService := fields.CreateService(s.helper)
	if len(item.Fields) > 0 {
		err = fieldsService.UpsertMany(item.Fields)
		if err != nil {
			return err
		}
	}
	if len(item.FieldsForDelete) > 0 {
		err = fieldsService.DeleteMany(item.FieldsForDelete)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
