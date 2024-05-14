package donorrules

import (
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.DonorRules, error) {
	return s.repository.getAll()
}

func (s *Service) UpsertMany(items RulesWithDeleted) error {
	err := fileinfos.CreateService(s.helper).UpsertMany(items.DonorRules.GetImages())
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
