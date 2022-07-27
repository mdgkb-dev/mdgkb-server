package visitsapplications

import (
	"mdgkb/mdgkb-server/handlers/formvalues"
	"mdgkb/mdgkb-server/handlers/visits"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.VisitsApplicationsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.VisitsApplication, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.VisitsApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = visits.CreateService(s.helper).UpsertMany(item.Visits)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.VisitsApplication) error {
	err := formvalues.CreateService(s.helper).Upsert(item.FormValue)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	err = visits.CreateService(s.helper).UpsertMany(item.Visits)
	if err != nil {
		return err
	}
	if len(item.VisitsForDelete) > 0 {
		err = visits.CreateService(s.helper).DeleteMany(item.VisitsForDelete)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) error {
	return s.repository.SetQueryFilter(c)
}
