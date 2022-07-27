package newsslides

import (
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/newsslidebuttons"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.NewsSlide) error {
	err := fileinfos.CreateService(s.helper).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = newsslidebuttons.CreateService(s.helper).CreateMany(item.NewsSlideButtons)
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
	err := fileinfos.CreateService(s.helper).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	newsSlideButtonsService := newsslidebuttons.CreateService(s.helper)
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
