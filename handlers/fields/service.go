package fields

import (
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/maskTokens"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.Field) error {
	item.SetForeignKeys()
	if item == nil {
		return nil
	}
	return s.repository.create(item)
}

func (s *Service) Update(item *models.Field) error {
	item.SetForeignKeys()
	if item == nil {
		return nil
	}
	return s.repository.update(item)
}

func (s *Service) UpsertMany(items models.Fields) error {
	if len(items) == 0 {
		return nil
	}
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(items.GetFileInfos())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	if len(items) == 0 {
		return nil
	}
	err = s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	maskTokensService := maskTokens.CreateService(s.repository.getDB())
	if len(items.GetMaskTokens()) > 0 {
		err = maskTokensService.UpsertMany(items.GetMaskTokens())
		if err != nil {
			return err
		}
	}
	if len(items.GetMaskTokensForDelete()) > 0 {
		err = maskTokensService.DeleteMany(items.GetMaskTokensForDelete())
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) Upsert(item *models.Field) error {
	item.SetForeignKeys()
	if item == nil {
		return nil
	}
	err := s.repository.upsert(item)
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
