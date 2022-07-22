package appointments

import (
	"mdgkb/mdgkb-server/handlers/children"
	"mdgkb/mdgkb-server/handlers/doctors"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.Appointments, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.Appointment, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.Appointment) error {
	err := children.CreateService(s.helper).Upsert(item.Child)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Appointment) error {
	err := children.CreateService(s.helper).Upsert(item.Child)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(items models.Appointments) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(id []string) error {
	if len(id) == 0 {
		return nil
	}
	return s.repository.deleteMany(id)
}

func (s *Service) Init() error {
	doctorsWithTimetable, err := doctors.CreateService(s.helper).GetAllTimetables()
	if err != nil {
		return err
	}
	doctorsWithTimetable.InitAppointmentsSlots()
	days := s.helper.Util.GetMonthDays()
	appointmentsToInsert := doctorsWithTimetable.InitAppointments(days)
	err = s.repository.upsertMany(appointmentsToInsert)
	if err != nil {
		return err
	}
	return nil
}
