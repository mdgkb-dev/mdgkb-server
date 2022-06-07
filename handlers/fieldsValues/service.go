package fieldsValues

import (
	"mdgkb/mdgkb-server/handlers/fieldValuesFiles"
	"mdgkb/mdgkb-server/handlers/fields"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.FieldValue) error {
	if item == nil {
		return nil
	}
	return s.repository.create(item)
}

func (s *Service) Update(item *models.FieldValue) error {
	if item == nil {
		return nil
	}
	return s.repository.update(item)
}

func (s *Service) UpsertMany(items models.FieldValues) error {
	if len(items) == 0 {
		return nil
	}
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(items.GetFileInfos())
	if err != nil {
		return err
	}
	err = fields.CreateService(s.repository.getDB()).UpsertMany(items.GetFields())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	fieldValuesFilesService := fieldValuesFiles.CreateService(s.repository.getDB())
	err = fieldValuesFilesService.UpsertMany(items.GetFieldValuesFiles())
	if err != nil {
		return err
	}
	err = fieldValuesFilesService.DeleteMany(items.GetFieldValuesFilesForDelete())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Upsert(item *models.FieldValue) error {
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
