package preparationrulesgroups

import (
	"mdgkb/mdgkb-server/handlers/preparationrules"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.PreparationRulesGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	err = preparationrules.CreateService(s.helper).CreateMany(items.GetPreparationRules())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.PreparationRulesGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	preparationRulesGroupsService := preparationrules.CreateService(s.helper)
	err = preparationRulesGroupsService.UpsertMany(items.GetPreparationRules())
	if err != nil {
		return err
	}
	err = preparationRulesGroupsService.DeleteMany(items.GetPreparationRulesForDelete())
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
