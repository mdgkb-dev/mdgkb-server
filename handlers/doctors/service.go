package doctors

import (
	"context"

	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/doctorpaidservices"
	"mdgkb/mdgkb-server/handlers/doctorsdivisions"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(c context.Context, item *models.Doctor) error {
	if item == nil {
		return nil
	}
	err := timetables.CreateService(s.helper).Create(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	if err != nil {
		return err
	}
	err = doctorpaidservices.CreateService(s.helper).CreateMany(item.DoctorPaidServices)
	if err != nil {
		return err
	}
	err = doctorsdivisions.CreateService(s.helper).CreateMany(item.DoctorsDivisions)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.Doctor) error {
	if item == nil {
		return nil
	}
	err := timetables.CreateService(s.helper).Upsert(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Upsert(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	doctorPaidServicesService := doctorpaidservices.CreateService(s.helper)
	err = doctorPaidServicesService.UpsertMany(item.DoctorPaidServices)
	if err != nil {
		return err
	}
	err = doctorPaidServicesService.DeleteMany(item.DoctorPaidServicesForDelete)
	if err != nil {
		return err
	}
	doctorsDivisionsService := doctorsdivisions.CreateService(s.helper)
	err = doctorsDivisionsService.UpsertMany(item.DoctorsDivisions)
	if err != nil {
		return err
	}
	err = doctorsDivisionsService.DeleteMany(item.DoctorsDivisionsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.Doctors, error) {
	return R.GetAll(c)
}

func (s *Service) GetAllTimetables(c context.Context) (models.Doctors, error) {
	return R.GetAllTimetables(c)
}

func (s *Service) Get(c context.Context, slug string) (*models.Doctor, error) {
	item, err := R.Get(c, slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetByDivisionID(c context.Context, divisionID string) (models.Doctors, error) {
	return R.GetByDivisionID(c, divisionID)
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) CreateComment(c context.Context, item *models.DoctorComment) error {
	commentsService := comments.S
	err := commentsService.UpsertOne(context.TODO(), item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return R.CreateComment(c, item)
}

func (s *Service) UpdateComment(c context.Context, item *models.DoctorComment) error {
	commentsService := comments.S
	err := commentsService.UpdateOne(context.TODO(), item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return R.UpdateComment(c, item)
}

func (s *Service) RemoveComment(c context.Context, id string) error {
	return R.RemoveComment(c, id)
}

func (s *Service) UpsertMany(c context.Context, items models.Doctors) error {
	if len(items) == 0 {
		return nil
	}
	return R.UpsertMany(c, items)
}

func (s *Service) CreateSlugs(c context.Context) error {
	_, err := R.GetAll(c)
	if err != nil {
		return err
	}
	humans := make(models.Humans, 0)
	//for i := range items {
	//	items[i].Human.Slug = s.helper.Util.MakeSlug(items[i].Human.GetFullName())
	//	humans = append(humans, items[i].Human)
	//}
	err = human.CreateService(s.helper).UpsertMany(humans)
	return err
}

func (s *Service) DeleteMany(c context.Context, idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return R.DeleteMany(c, idPool)
}
