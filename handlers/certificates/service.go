package certificates

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.Certificates, error) {
	return s.repository.getAll()
}

func (s *Service) CreateMany(items models.Certificates) error {
	if len(items) == 0 {
		return nil
	}
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(items.GetFileInfos())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.createMany(items)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpsertMany(items models.Certificates) error {
	if len(items) == 0 {
		return nil
	}
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(items.GetFileInfos())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.upsertMany(items)
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
