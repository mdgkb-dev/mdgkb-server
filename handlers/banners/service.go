package banners

import (
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Banner) error {
	err := fileInfos.CreateService(s.helper).Create(item.FileInfo)
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

func (s *Service) Update(item *models.Banner) error {
	err := fileInfos.CreateService(s.helper).Upsert(item.FileInfo)
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

func (s *Service) GetAll() (models.Banners, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.Banner, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) UpdateAllOrder(items models.Banners) error {
	err := s.repository.updateAllOrder(items)
	if err != nil {
		return err
	}
	return nil
}
