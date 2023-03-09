package residencycoursepracticeplacegroups

import (
	"mdgkb/mdgkb-server/handlers/residencycoursepracticeplaces"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.ResidencyCoursePracticePlaceGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	err = residencycoursepracticeplaces.CreateService(s.helper).CreateMany(items.GetResidencyCoursePracticePlaces())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.ResidencyCoursePracticePlaceGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	residencyCoursePracticePlacesService := residencycoursepracticeplaces.CreateService(s.helper)
	err = residencyCoursePracticePlacesService.UpsertMany(items.GetResidencyCoursePracticePlaces())
	if err != nil {
		return err
	}
	err = residencyCoursePracticePlacesService.DeleteMany(items.GetResidencyCoursePracticePlacesForDelete())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
