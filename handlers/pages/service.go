package pages

import (
	"context"

	"mdgkb/mdgkb-server/handlers/pageimages"
	"mdgkb/mdgkb-server/handlers/pagesdocuments"
	"mdgkb/mdgkb-server/handlers/pagesidemenus"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(c context.Context, item *models.Page) error {
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	item.SetForeignKeys()
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	pageSideMenusService := pagesidemenus.CreateService(s.helper)
	err = pageSideMenusService.UpsertMany(item.PageSideMenus)
	if err != nil {
		return err
	}

	pagesDocumentsService := pagesdocuments.CreateService(s.helper)
	err = pagesDocumentsService.UpsertMany(item.PageDocuments)
	if err != nil {
		return err
	}
	err = pagesDocumentsService.DeleteMany(item.PageDocumentsForDelete)
	if err != nil {
		return err
	}
	pageImagesService := pageimages.CreateService(s.helper)
	err = pageImagesService.UpsertMany(item.PageImages)
	if err != nil {
		return err
	}
	err = pageImagesService.DeleteMany(item.PageImagesForDelete)
	if err != nil {
		return err
	}
	return nil
	//pagesCommentsService := .CreateService(s.helper)
	//err = pagesCommentsService.CreateMany(item.PageComments)
	//if err != nil {
	//	return err
	//}
}

func (s *Service) GetAll(c context.Context) (models.PagesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id *string) (*models.Page, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.Page) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	pageSideMenusService := pagesidemenus.CreateService(s.helper)
	err = pageSideMenusService.UpsertMany(item.PageSideMenus)
	if err != nil {
		return err
	}
	err = pageSideMenusService.DeleteMany(item.PageSideMenusForDelete)
	if err != nil {
		return err
	}

	pagesDocumentsService := pagesdocuments.CreateService(s.helper)
	err = pagesDocumentsService.UpsertMany(item.PageDocuments)
	if err != nil {
		return err
	}
	err = pagesDocumentsService.DeleteMany(item.PageDocumentsForDelete)
	if err != nil {
		return err
	}
	pageImagesService := pageimages.CreateService(s.helper)
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

func (s *Service) Delete(c context.Context, id *string) error {
	return R.Delete(c, id)
}

func (s *Service) GetBySlug(c context.Context, slug *string) (*models.Page, error) {
	item, err := R.GetBySlug(c, slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}
