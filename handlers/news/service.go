package news

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/events"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/newsDivisions"
	"mdgkb/mdgkb-server/handlers/newsDoctors"
	"mdgkb/mdgkb-server/handlers/newsImages"
	"mdgkb/mdgkb-server/handlers/newsToCategories"
	"mdgkb/mdgkb-server/handlers/newsToTags"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.News) error {
	err := fileInfos.CreateService(s.helper).UpsertMany(item.GetFileInfos())
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
	item.SetIdForChildren()

	err = newsToCategories.CreateService(s.helper).CreateMany(item.NewsToCategories)
	if err != nil {
		return err
	}
	err = newsToTags.CreateService(s.helper).CreateMany(item.NewsToTags)
	if err != nil {
		return err
	}
	err = newsImages.CreateService(s.helper).CreateMany(item.NewsImages)
	if err != nil {
		return err
	}
	err = newsDoctors.CreateService(s.helper).CreateMany(item.NewsDoctors)
	if err != nil {
		return err
	}
	err = newsDivisions.CreateService(s.helper).CreateMany(item.NewsDivisions)
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
	err := fileInfos.CreateService(s.helper).UpsertMany(item.GetFileInfos())
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
	item.SetIdForChildren()

	err = newsToCategories.CreateService(s.helper).UpsertMany(item.NewsToCategories)
	if err != nil {
		return err
	}
	err = newsToTags.CreateService(s.helper).UpsertMany(item.NewsToTags)
	if err != nil {
		return err
	}
	newsImagersService := newsImages.CreateService(s.helper)
	err = newsImagersService.UpsertMany(item.NewsImages)
	if err != nil {
		return err
	}
	err = newsImagersService.DeleteMany(item.NewsImagesForDelete)
	if err != nil {
		return err
	}
	newsDoctorsService := newsDoctors.CreateService(s.helper)
	err = newsDoctorsService.UpsertMany(item.NewsDoctors)
	if err != nil {
		return err
	}
	err = newsDoctorsService.DeleteMany(item.NewsDoctorsForDelete)
	if err != nil {
		return err
	}

	newsDivisionsService := newsDivisions.CreateService(s.helper)
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
