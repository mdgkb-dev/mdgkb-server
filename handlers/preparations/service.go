package preparations

import (
	"mdgkb/mdgkb-server/handlers/preparationrulesgroups"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Preparation) error {
	err := s.repository.Create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = preparationrulesgroups.CreateService(s.helper).CreateMany(item.PreparationRulesGroups)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Preparation) error {
	err := s.repository.Update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	preparationRulesGroupsService := preparationrulesgroups.CreateService(s.helper)
	err = preparationRulesGroupsService.UpsertMany(item.PreparationRulesGroups)
	if err != nil {
		return err
	}
	err = preparationRulesGroupsService.DeleteMany(item.PreparationRulesGroupsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.PreparationsWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(id string) (*models.Preparation, error) {
	item, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
