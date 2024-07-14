package heads

import (
	"context"
	"mdgkb/mdgkb-server/handlers/contactinfo"
	"mdgkb/mdgkb-server/handlers/departments"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(c context.Context, item *models.Head) error {
	if item == nil {
		return nil
	}
	err := contactinfo.CreateService(s.helper).Create(item.ContactInfo)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.helper).Create(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Upsert(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = departments.CreateService(s.helper).CreateMany(item.Departments)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.Head) error {
	if item == nil {
		return nil
	}
	err := contactinfo.CreateService(s.helper).Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.helper).Upsert(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Upsert(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	departmentsService := departments.CreateService(s.helper)
	err = departmentsService.UpsertMany(item.Departments)
	if err != nil {
		return err
	}
	err = departmentsService.DeleteMany(item.DepartmentsForDelete)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll(c context.Context) (models.Heads, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Head, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) UpdateAll(c context.Context, items models.Heads) error {
	return R.UpdateAll(c, items)
}

func (s *Service) DeleteMany(c context.Context, idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return R.DeleteMany(c, idPool)
}
