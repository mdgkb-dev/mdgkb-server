package news

import (
	"context"
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

func (s *Service) Create(c context.Context, item *models.News) error {
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
	err = R.Create(c, item)
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

func (s *Service) GetAll(c context.Context) (models.NewsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) GetMain(c context.Context) (models.NewsWithCount, error) {
	return R.GetMain(c)
}

func (s *Service) GetSubMain(c context.Context) (models.NewsWithCount, error) {
	return R.GetSubMain(c)
}

func (s *Service) RemoveTag(c context.Context, item *models.NewsToTag) error {
	return R.RemoveTag(c, item)
}

func (s *Service) AddTag(c context.Context, item *models.NewsToTag) error {
	return R.RemoveTag(c, item)
}

func (s *Service) CreateLike(c context.Context, item *models.NewsLike) error {
	return R.CreateLike(c, item)
}

func (s *Service) CreateComment(c context.Context, item *models.NewsComment) error {
	commentsService := comments.S
	err := commentsService.UpsertOne(context.TODO(), item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return R.CreateComment(c, item)
}

func (s *Service) UpdateComment(c context.Context, item *models.NewsComment) error {
	commentsService := comments.S
	err := commentsService.UpdateOne(context.TODO(), item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return R.UpdateComment(c, item)
}

func (s *Service) CreateViewOfNews(c context.Context, item *models.NewsView) error {
	return R.CreateViewOfNews(c, item)
}

func (s *Service) RemoveComment(c context.Context, id string) error {
	return R.RemoveComment(c, id)
}

func (s *Service) DeleteLike(c context.Context, id string) error {
	return R.DeleteLike(c, id)
}

func (s *Service) GetBySlug(c context.Context, slug string) (*models.News, error) {
	item, err := R.GetBySlug(c, slug)
	if err != nil {
		return nil, err
	}
	item.ViewsCount = len(item.NewsViews)

	// g := &models.GeoIP{}
	// country, city := g.GetByIP(c.ClientIP())

	// newsView := models.NewsView{IPAddress: c.ClientIP(), NewsID: item.ID, Country: country, City: city}
	// err = s.CreateViewOfNews(&newsView)
	// if err != nil {
	// 	return nil, err
	// }
	// if newsView.ID.Valid {
	// 	item.ViewsCount++
	// }
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.News) error {
	err := fileinfos.CreateService(s.helper).UpsertMany(item.GetFileInfos())
	if err != nil {
		return err
	}
	err = events.CreateService(s.helper).Upsert(item.Event)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Update(c, item)
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

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) GetSuggestionNews(c context.Context, id string) ([]*models.News, error) {
	return R.GetSuggestionNews(c, id)
}

func (s *Service) GetAggregateViews(c context.Context, opts *exportmodels.NewsView) (models.ChartDataSets, error) {
	return R.GetAggregateViews(c, opts)
}

func (s *Service) GetNewsComments(c *gin.Context, id string) (*models.NewsComments, error) {
	return R.GetNewsComments(c, id)
}
