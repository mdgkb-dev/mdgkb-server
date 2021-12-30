package preparationRulesGroups

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/preparationRules"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.PreparationRulesGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	err = preparationRules.CreateService(s.repository.getDB()).CreateMany(items.GetPreparationRules())
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
	items.SetIdForChildren()
	preparationRulesGroupsService := preparationRules.CreateService(s.repository.getDB())
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
