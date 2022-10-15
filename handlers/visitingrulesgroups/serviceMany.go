package visitingrulesgroups

import (
	"mdgkb/mdgkb-server/handlers/visitingrules"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.VisitingRulesGroups) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items models.VisitingRulesGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	visitingRulesService := visitingrules.CreateService(s.helper)
	err = visitingRulesService.UpsertMany(items.GetVisitingRules())
	if err != nil {
		return err
	}
	err = visitingRulesService.DeleteMany(items.GetVisitingRulesForDelete())
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
