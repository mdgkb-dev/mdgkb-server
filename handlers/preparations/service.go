package preparations

import (
	"mdgkb/mdgkb-server/handlers/preparationRulesGroups"
	"mdgkb/mdgkb-server/handlers/preparationsToTags"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Preparation) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = preparationRulesGroups.CreateService(s.helper).CreateMany(item.PreparationRulesGroups)
	if err != nil {
		return err
	}

	err = preparationsToTags.CreateService(s.helper).CreateMany(item.PreparationsToTags)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Preparation) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	preparationRulesGroupsService := preparationRulesGroups.CreateService(s.helper)
	err = preparationRulesGroupsService.UpsertMany(item.PreparationRulesGroups)
	if err != nil {
		return err
	}
	err = preparationRulesGroupsService.DeleteMany(item.PreparationRulesGroupsForDelete)
	if err != nil {
		return err
	}

	preparationsToTagsService := preparationsToTags.CreateService(s.helper)
	err = preparationsToTagsService.UpsertMany(item.PreparationsToTags)
	if err != nil {
		return err
	}
	err = preparationsToTagsService.DeleteMany(item.PreparationsToTagsForDelete)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll() (models.Preparations, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.Preparation, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(item PreparationsWithDeleted) error {
	err := s.repository.upsertMany(item.Preparations)
	if err != nil {
		return err
	}
	if len(item.PreparationsForDeleted) > 0 {
		err = s.repository.deleteMany(item.PreparationsForDeleted)
		if err != nil {
			return err
		}
	}

	item.Preparations.SetIdForChildren()
	preparationRulesGroupsService := preparationRulesGroups.CreateService(s.helper)
	err = preparationRulesGroupsService.UpsertMany(item.Preparations.GetPreparationRulesGroups())
	if err != nil {
		return err
	}
	err = preparationRulesGroupsService.DeleteMany(item.Preparations.GetPreparationRulesGroupsForDelete())
	if err != nil {
		return err
	}

	preparationsToTagsService := preparationsToTags.CreateService(s.helper)
	err = preparationsToTagsService.UpsertMany(item.Preparations.GetPreparationsToTags())
	if err != nil {
		return err
	}
	err = preparationsToTagsService.DeleteMany(item.Preparations.GetPreparationsToTagsForDelete())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetTags() (models.PreparationsTags, error) {
	return s.repository.getTags()
}
