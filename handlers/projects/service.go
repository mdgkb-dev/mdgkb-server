package projects

import (
	"context"
	"mdgkb/mdgkb-server/handlers/projectitems"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(c context.Context, item *models.Project) error {
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	projectItemsService := projectitems.CreateService(s.helper)
	err = projectItemsService.UpsertMany(item.ProjectItems)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.Project) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	projectItemsService := projectitems.CreateService(s.helper)
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

func (s *Service) GetAll(c context.Context) (models.Projects, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Project, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) GetBySlug(c context.Context, slug string) (*models.Project, error) {
	item, err := R.GetBySlug(c, slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}
