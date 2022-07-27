package news

import (
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/events"
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/newsdivisions"
	"mdgkb/mdgkb-server/handlers/newsdoctors"
	"mdgkb/mdgkb-server/handlers/newsimages"
	"mdgkb/mdgkb-server/handlers/newstocategories"
	"mdgkb/mdgkb-server/handlers/newstotags"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.News) error {
	err := fileinfos.CreateService(s.helper).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	err = events.CreateService(s.helper).Create(item.Event)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.Title, true)
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = newstocategories.CreateService(s.helper).CreateMany(item.NewsToCategories)
	if err != nil {
		return err
	}
	err = newstotags.CreateService(s.helper).CreateMany(item.NewsToTags)
	if err != nil {
		return err
	}
	err = newsimages.CreateService(s.helper).CreateMany(item.NewsImages)
	if err != nil {
		return err
	}
	err = newsdoctors.CreateService(s.helper).CreateMany(item.NewsDoctors)
	if err != nil {
		return err
	}
	err = newsdivisions.CreateService(s.helper).CreateMany(item.NewsDivisions)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.NewsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) RemoveTag(item *models.NewsToTag) error {
	return s.repository.removeTag(item)
}

func (s *Service) AddTag(item *models.NewsToTag) error {
	return s.repository.removeTag(item)
}

func (s *Service) CreateLike(item *models.NewsLike) error {
	return s.repository.createLike(item)
}

func (s *Service) CreateComment(item *models.NewsComment) error {
	commentsService := comments.CreateService(s.helper)
	err := commentsService.UpsertOne(item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return s.repository.createComment(item)
}

func (s *Service) UpdateComment(item *models.NewsComment) error {
	commentsService := comments.CreateService(s.helper)
	err := commentsService.UpdateOne(item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return s.repository.updateComment(item)
}

func (s *Service) CreateViewOfNews(item *models.NewsView) error {
	return s.repository.createViewOfNews(item)
}

func (s *Service) RemoveComment(id string) error {
	return s.repository.removeComment(id)
}

func (s *Service) DeleteLike(id string) error {
	return s.repository.deleteLike(id)
}

func (s *Service) GetBySlug(slug string) (*models.News, error) {
	return s.repository.getBySlug(slug)
}

func (s *Service) Update(item *models.News) error {
	err := fileinfos.CreateService(s.helper).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	err = events.CreateService(s.helper).Upsert(item.Event)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = newstocategories.CreateService(s.helper).UpsertMany(item.NewsToCategories)
	if err != nil {
		return err
	}
	err = newstotags.CreateService(s.helper).UpsertMany(item.NewsToTags)
	if err != nil {
		return err
	}
	newsImagersService := newsimages.CreateService(s.helper)
	err = newsImagersService.UpsertMany(item.NewsImages)
	if err != nil {
		return err
	}
	err = newsImagersService.DeleteMany(item.NewsImagesForDelete)
	if err != nil {
		return err
	}
	newsDoctorsService := newsdoctors.CreateService(s.helper)
	err = newsDoctorsService.UpsertMany(item.NewsDoctors)
	if err != nil {
		return err
	}
	err = newsDoctorsService.DeleteMany(item.NewsDoctorsForDelete)
	if err != nil {
		return err
	}

	newsDivisionsService := newsdivisions.CreateService(s.helper)
	err = newsDivisionsService.UpsertMany(item.NewsDivisions)
	if err != nil {
		return err
	}
	err = newsDivisionsService.DeleteMany(item.NewsDivisionsForDelete)
	if err != nil {
		return err
	}
	return err
}
func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) error {
	return s.repository.SetQueryFilter(c)
}
