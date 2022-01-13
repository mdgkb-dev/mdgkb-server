package news

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/handlers/events"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/newsImages"
	"mdgkb/mdgkb-server/handlers/newsToCategories"
	"mdgkb/mdgkb-server/handlers/newsToTags"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.News) error {
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	err = events.CreateService(s.repository.getDB(), s.helper).Create(item.Event)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	item.Slug = s.helper.MakeSlug(item.Title)
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = newsToCategories.CreateService(s.repository.getDB()).CreateMany(item.NewsToCategories)
	if err != nil {
		return err
	}
	err = newsToTags.CreateService(s.repository.getDB()).CreateMany(item.NewsToTags)
	if err != nil {
		return err
	}
	err = newsImages.CreateService(s.repository.getDB()).CreateMany(item.NewsImages)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll(params *newsParams) ([]models.News, error) {
	return s.repository.getAll(params)
}

func (s *Service) GetAllAdmin() (models.NewsWithCount, error) {
	return s.repository.getAllAdmin()
}

func (s *Service) GetAllRelationsNews(params *newsParams) ([]models.News, error) {
	return s.repository.getAllRelationsNews(params)
}

func (s *Service) GetByMonth(params *monthParams) ([]models.News, error) {
	return s.repository.getByMonth(params)
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
	return s.repository.createComment(item)
}

func (s *Service) UpdateComment(item *models.NewsComment) error {
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
	err := fileInfos.CreateService(s.repository.getDB()).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	err = events.CreateService(s.repository.getDB(), s.helper).Upsert(item.Event)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = newsToCategories.CreateService(s.repository.getDB()).UpsertMany(item.NewsToCategories)
	if err != nil {
		return err
	}
	err = newsToTags.CreateService(s.repository.getDB()).UpsertMany(item.NewsToTags)
	if err != nil {
		return err
	}
	err = newsImages.CreateService(s.repository.getDB()).UpsertMany(item.NewsImages)
	if err != nil {
		return err
	}
	return err
}
func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
