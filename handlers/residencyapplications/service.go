package residencyapplications

import (
	"mdgkb/mdgkb-server/handlers/diplomas"
	"mdgkb/mdgkb-server/handlers/formvalues"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/handlers/residencyapplicationspointsachievements"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.ResidencyApplicationsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.ResidencyApplication, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) EmailExists(email string, courseID string) (bool, error) {
	item, err := s.repository.emailExists(email, courseID)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) TypeExists(email string, main bool) (bool, error) {
	item, err := s.repository.typeExists(email, main)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) Create(item *models.ResidencyApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	err = diplomas.CreateService(s.helper).Upsert(item.Diploma)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	residencyApplicationsPointsAchievementsService := residencyapplicationspointsachievements.CreateService(s.helper)
	err = residencyApplicationsPointsAchievementsService.CreateMany(item.ResidencyApplicationPointsAchievements)
	if err != nil {
		return err
	}
	err = residencyApplicationsPointsAchievementsService.DeleteMany(item.ResidencyApplicationPointsAchievementsForDelete)
	if err != nil {
		return err
	}
	err = meta.CreateService(s.helper).SendApplicationsCounts()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.ResidencyApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	err = diplomas.CreateService(s.helper).Upsert(item.Diploma)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	residencyApplicationsPointsAchievementsService := residencyapplicationspointsachievements.CreateService(s.helper)
	err = residencyApplicationsPointsAchievementsService.UpsertMany(item.ResidencyApplicationPointsAchievements)
	if err != nil {
		return err
	}
	err = residencyApplicationsPointsAchievementsService.DeleteMany(item.ResidencyApplicationPointsAchievementsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateWithForm(item *models.FormValue) error {
	err := formvalues.CreateService(s.helper).Upsert(item)
	if err != nil {
		return err
	}
	err = s.repository.update(item.ResidencyApplication)
	if err != nil {
		return err
	}
	item.ResidencyApplication.SetIDForChildren()
	residencyApplicationsPointsAchievementsService := residencyapplicationspointsachievements.CreateService(s.helper)
	err = residencyApplicationsPointsAchievementsService.UpsertMany(item.ResidencyApplication.ResidencyApplicationPointsAchievements)
	if err != nil {
		return err
	}
	err = residencyApplicationsPointsAchievementsService.DeleteMany(item.ResidencyApplication.ResidencyApplicationPointsAchievementsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.ResidencyApplications) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
