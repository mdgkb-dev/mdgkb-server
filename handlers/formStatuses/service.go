package formStatuses

import (
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/formStatusToFormStatuses"
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

func (s *Service) GetAllByGroupId(id *string) (models.FormStatuses, error) {
	return s.repository.GetAllByGroupId(id)
}

func (s *Service) Get(id *string) (*models.FormStatus, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Upsert(item *models.FormStatus) error {
	err := fileInfos.CreateService(s.repository.getDB()).Upsert(item.Icon)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = formStatusToFormStatuses.CreateService(s.repository.getDB()).UpsertMany(item.FormStatusToFormStatuses)
	if err != nil {
		return err
	}
	err = formStatusToFormStatuses.CreateService(s.repository.getDB()).DeleteMany(item.FormStatusToFormStatusesForDelete)
	return nil
}

func (s *Service) UpsertMany(items models.FormStatuses) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	err = formStatusToFormStatuses.CreateService(s.repository.getDB()).UpsertMany(items.GetFormStatusToFormStatuses())
	if err != nil {
		return err
	}
	if len(items.GetFormStatusToFormStatusesForDelete()) > 0 {
		err = formStatusToFormStatuses.CreateService(s.repository.getDB()).DeleteMany(items.GetFormStatusToFormStatusesForDelete())
		if err != nil {
			return err
		}
	}
	return nil

}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
