package paidprogramsgroups

import (
	"mdgkb/mdgkb-server/handlers/paidprograms"
	"mdgkb/mdgkb-server/handlers/paidprogramservices"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.PaidProgramsGroup) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = paidprograms.CreateService(s.helper).CreateMany(item.PaidPrograms)
	if err != nil {
		return err
	}

	err = paidprogramservices.CreateService(s.helper).CreateMany(item.PaidProgramServices)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.PaidProgramsGroup) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	paidProgramsService := paidprograms.CreateService(s.helper)
	err = paidProgramsService.UpsertMany(item.PaidPrograms)
	if err != nil {
		return err
	}
	err = paidProgramsService.DeleteMany(item.PaidProgramServicesForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.PaidProgramsGroups, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.PaidProgramsGroup, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(item models.PaidProgramsGroupsStruct) error {
	err := s.repository.upsertMany(item.PaidProgramsGroups)
	if err != nil {
		return err
	}
	if len(item.PaidProgramsGroupsForDelete) > 0 {
		err = s.repository.deleteMany(item.PaidProgramsGroupsForDelete)
		if err != nil {
			return err
		}
	}

	item.PaidProgramsGroups.SetIDForChildren()
	paidProgramsService := paidprograms.CreateService(s.helper)
	err = paidProgramsService.UpsertMany(item.PaidProgramsGroups.GetPaidPrograms())
	if err != nil {
		return err
	}
	err = paidProgramsService.DeleteMany(item.PaidProgramsGroups.GetPaidProgramsForDelete())
	if err != nil {
		return err
	}
	return nil
}
