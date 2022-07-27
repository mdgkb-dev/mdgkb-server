package formstatuses

import (
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/formstatustoformstatuses"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) GetAll() (models.FormStatuses, error) {
	return s.repository.getAll()
}

func (s *Service) GetAllByGroupID(id *string) (models.FormStatuses, error) {
	return s.repository.GetAllByGroupID(id)
}

func (s *Service) Get(id *string) (*models.FormStatus, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Upsert(item *models.FormStatus) error {
	err := fileinfos.CreateService(s.helper).Upsert(item.Icon)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = formstatustoformstatuses.CreateService(s.helper).UpsertMany(item.FormStatusToFormStatuses)
	if err != nil {
		return err
	}
	err = formstatustoformstatuses.CreateService(s.helper).DeleteMany(item.FormStatusToFormStatusesForDelete)
	return err
}

func (s *Service) UpsertMany(items models.FormStatuses) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	err = formstatustoformstatuses.CreateService(s.helper).UpsertMany(items.GetFormStatusToFormStatuses())
	if err != nil {
		return err
	}
	if len(items.GetFormStatusToFormStatusesForDelete()) > 0 {
		err = formstatustoformstatuses.CreateService(s.helper).DeleteMany(items.GetFormStatusToFormStatusesForDelete())
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
