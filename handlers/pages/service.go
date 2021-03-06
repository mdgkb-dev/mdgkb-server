package pages

import (
	"mdgkb/mdgkb-server/handlers/pageImages"
	"mdgkb/mdgkb-server/handlers/pagesDocuments"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Page) error {
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = pagesDocuments.CreateService(s.repository.getDB()).CreateMany(item.PageDocuments)
	if err != nil {
		return err
	}
	err = pageImages.CreateService(s.repository.getDB()).CreateMany(item.PageImages)
	if err != nil {
		return err
	}
	//pagesCommentsService := .CreateService(s.repository.getDB())
	//err = pagesCommentsService.CreateMany(item.PageComments)
	//if err != nil {
	//	return err
	//}
	return err
}

func (s *Service) GetAll() (models.Pages, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.Page, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Page) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	pagesDocumentsService := pagesDocuments.CreateService(s.repository.getDB())
	err = pagesDocumentsService.UpsertMany(item.PageDocuments)
	if err != nil {
		return err
	}
	err = pagesDocumentsService.DeleteMany(item.PageDocumentsForDelete)
	if err != nil {
		return err
	}
	pageImagesService := pageImages.CreateService(s.repository.getDB())
	err = pageImagesService.UpsertMany(item.PageImages)
	if err != nil {
		return err
	}
	err = pageImagesService.DeleteMany(item.PageImagesForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) GetBySlug(slug *string) (*models.Page, error) {
	item, err := s.repository.getBySlug(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}
