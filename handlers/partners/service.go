package partners

import (
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.Partners, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.Partner, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.Partner) error {
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Partner) error {
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
