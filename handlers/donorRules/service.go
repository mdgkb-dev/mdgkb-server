package donorRules

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll(userId *uuid.UUID) (models.DonorRules, error) {
	return s.repository.getAll(userId)
}

func (s *Service) UpsertMany(items DonorRulesWithDeleted) error {
	err := fileInfos.CreateService(s.helper).UpsertMany(items.DonorRules.GetImages())
	if err != nil {
		return err
	}
	items.DonorRules.SetForeignKeys()
	err = s.repository.upsertMany(items.DonorRules)
	if err != nil {
		return err
	}
	if len(items.DonorRulesForDelete) > 0 {
		err = s.repository.deleteMany(items.DonorRulesForDelete)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) AddToUser(donorRuleUser *models.DonorRuleUser) error {
	return s.repository.addToUser(donorRuleUser)
}

func (s *Service) DeleteFromUser(donorRuleUser *models.DonorRuleUser) error {
	return s.repository.deleteFromUser(donorRuleUser)
}
