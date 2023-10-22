package news

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	//_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(news *models.News) (err error) {
	_, err = r.DB().NewInsert().Model(news).Exec(r.ctx)
	return err
}

func (r *Repository) update(news *models.News) (err error) {
	_, err = r.DB().NewUpdate().Model(news).Where("id = ?", news.ID).Exec(r.ctx)
	return err
}

func (r *Repository) createLike(item *models.NewsLike) error {
	_, err := r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) addTag(item *models.NewsToTag) error {
	_, err := r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) removeTag(item *models.NewsToTag) error {
	_, err := r.DB().NewDelete().Model(&models.NewsToTag{}).Where("news_id = ? AND tag_id = ?", item.NewsID, item.TagID).Exec(r.ctx)
	return err
}

func (r *Repository) createComment(item *models.NewsComment) error {
	_, err := r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) updateComment(item *models.NewsComment) error {
	_, err := r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) removeComment(id string) error {
	_, err := r.DB().NewDelete().Model(&models.NewsComment{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.NewsWithCount, err error) {
	items.News = make([]*models.News, 0)
	query := r.DB().NewSelect().Model(&items.News).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("PreviewImage").
		Relation("NewsLikes").
		Relation("NewsViews")
	r.queryFilter.HandleQuery(query)
	fmt.Println(time.Now())
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) getMain() (items models.NewsWithCount, err error) {
	items.News = make([]*models.News, 0)
	query := r.DB().NewSelect().Model(&items.News).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("PreviewImage").
		Relation("NewsLikes").
		Relation("NewsViews").
		Where("news_view.main = ?", true)
	r.queryFilter.HandleQuery(query)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) getSubMain() (items models.NewsWithCount, err error) {
	items.News = make([]*models.News, 0)
	query := r.DB().NewSelect().Model(&items.News).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("PreviewImage").
		Relation("NewsLikes").
		Relation("NewsViews").
		Where("news_view.sub_main = ?", true)
	r.queryFilter.HandleQuery(query)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) getBySlug(slug string) (*models.News, error) {
	item := models.News{}
	err := r.DB().NewSelect().Model(&item).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("PreviewImage").
		Relation("MainImage").
		Relation("NewsLikes").
		Relation("NewsViews").
		Relation("Event.Form.Fields.ValueType").
		Relation("Event.EventApplications.FieldValues").
		Relation("Event.EventApplications.User").
		Relation("NewsComments.Comment", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("comment.mod_checked = true").Order("comment.published_on DESC")
		}).
		Relation("NewsComments.Comment.User.Human").
		Relation("NewsImages", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("news_images.news_image_order")
		}).
		Relation("NewsImages.FileInfo").
		Relation("NewsDoctors.Doctor").
		Relation("NewsDoctors.Doctor.Employee.Human").
		Relation("NewsDoctors.Doctor.Employee.Human.PhotoMini").
		Relation("NewsDoctors.Doctor.Employee.Regalias").
		Where("news_view.slug = ?", slug).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.News{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) deleteLike(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.NewsLike{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) createViewOfNews(newsView *models.NewsView) (err error) {
	_, err = r.DB().NewInsert().Model(newsView).On("CONFLICT (ip_address, news_id) DO NOTHING").Exec(r.ctx)
	return err
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}
