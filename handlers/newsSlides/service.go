package newsSlides

import (
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/newsSlideButtons"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.NewsSlide) error {
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = newsSlideButtons.CreateService(s.repository.getDB()).CreateMany(item.NewsSlideButtons)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.NewsSlides, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.NewsSlide, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.NewsSlide) error {
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	newsSlideButtonsService := newsSlideButtons.CreateService(s.repository.getDB())
	err = newsSlideButtonsService.UpsertMany(item.NewsSlideButtons)
	if err != nil {
		return err
	}
	if len(item.NewsSlideButtonsForDelete) > 0 {
		err = newsSlideButtonsService.DeleteMany(item.NewsSlideButtonsForDelete)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) UpdateAll(items models.NewsSlides) error {
	return s.repository.updateAll(items)
}
