package projects

import (
	"mdgkb/mdgkb-server/handlers/projectItems"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Project) error {
	item.Slug = s.helper.Util.MakeSlug(item.Title)
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	projectItemsService := projectItems.CreateService(s.repository.getDB(), s.helper)
	err = projectItemsService.UpsertMany(item.ProjectItems)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Project) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	projectItemsService := projectItems.CreateService(s.repository.getDB(), s.helper)
	err = projectItemsService.UpsertMany(item.ProjectItems)
	if err != nil {
		return err
	}
	err = projectItemsService.DeleteMany(item.ProjectItemsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.Projects, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.Project, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) GetBySlug(slug *string) (*models.Project, error) {
	item, err := s.repository.getBySlug(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}
