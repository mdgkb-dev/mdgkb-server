package newsslides

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.NewsSlide) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.NewsSlides, error) {
	items := make(models.NewsSlides, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("NewsSlideButtons").
		Relation("DesktopImg").
		Relation("LaptopImg").
		Relation("TabletImg").
		Relation("MobileImg").
		Order("news_slide_order").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.NewsSlide, error) {
	item := models.NewsSlide{}
	err := r.db().NewSelect().Model(&item).
		Relation("NewsSlideButtons").
		Relation("DesktopImg").
		Relation("LaptopImg").
		Relation("TabletImg").
		Relation("MobileImg").
		Where("news_slides.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.NewsSlide{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.NewsSlide) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) updateAll(items models.NewsSlides) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("news_slide_order = EXCLUDED.news_slide_order").
		Where("news_slides.id = EXCLUDED.id").
		Exec(r.ctx)
	return err
}
