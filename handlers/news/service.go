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
	"mdgkb/mdgkb-server/models/exportmodels"

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
	err = s.Repository.create(item)
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

func (s *Service) GetAll(ftsp bool) (models.NewsWithCount, error) {
	return s.Repository.getAll(ftsp)
}

func (s *Service) GetMain() (models.NewsWithCount, error) {
	return s.Repository.getMain()
}

func (s *Service) GetSubMain() (models.NewsWithCount, error) {
	return s.Repository.getSubMain()
}

func (s *Service) RemoveTag(item *models.NewsToTag) error {
	return s.Repository.removeTag(item)
}

func (s *Service) AddTag(item *models.NewsToTag) error {
	return s.Repository.removeTag(item)
}

func (s *Service) CreateLike(item *models.NewsLike) error {
	return s.Repository.createLike(item)
}

func (s *Service) CreateComment(item *models.NewsComment) error {
	commentsService := comments.CreateService(s.helper)
	err := commentsService.UpsertOne(item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return s.Repository.createComment(item)
}

func (s *Service) UpdateComment(item *models.NewsComment) error {
	commentsService := comments.CreateService(s.helper)
	err := commentsService.UpdateOne(item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return s.Repository.updateComment(item)
}

func (s *Service) CreateViewOfNews(item *models.NewsView) error {
	return s.Repository.createViewOfNews(item)
}

func (s *Service) RemoveComment(id string) error {
	return s.Repository.removeComment(id)
}

func (s *Service) DeleteLike(id string) error {
	return s.Repository.deleteLike(id)
}

func (s *Service) GetBySlug(c *gin.Context, slug string) (*models.News, error) {
	item, err := s.Repository.getBySlug(slug)
	if err != nil {
		return nil, err
	}
	item.ViewsCount = len(item.NewsViews)

	g := &models.GeoIP{}
	country, city := g.GetByIP(c.ClientIP())

	newsView := models.NewsView{IPAddress: c.ClientIP(), NewsID: item.ID, Country: country, City: city}
	err = s.CreateViewOfNews(&newsView)
	if err != nil {
		return nil, err
	}
	if newsView.ID.Valid {
		item.ViewsCount++
	}
	return item, nil
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
	err = s.Repository.update(item)
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
	err = newstotags.CreateService(s.helper).DeleteMany(item.NewsToTagsForDelete)
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
	return s.Repository.delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context, item models.FTSPQuery) error {
	return s.Repository.SetQueryFilter(c, item)
}

func (s *Service) GetSuggestionNews(id string) ([]*models.News, error) {
	return s.Repository.GetSuggestionNews(id)
}

func (s *Service) GetAggregateViews(opts *exportmodels.NewsView) (models.ChartDataSets, error) {
	return s.Repository.GetAggregateViews(opts)
}

func (s *Service) GetNewsComments(c *gin.Context, id string) (*models.NewsComments, error) {
	return s.Repository.getNewsComments(id)
}
