package banners

import (
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getAll() (models.Banners, error) {
	items := make(models.Banners, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("FileInfo").
		Order("list_number").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Banner, error) {
	item := models.Banner{}
	err := r.db().NewSelect().Model(&item).Where("banners.id = ?", id).
		Relation("FileInfo").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.Banner) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Banner{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Banner) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) updateAllOrder(items models.Banners) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("list_number = EXCLUDED.list_number").
		Where("banner.id = EXCLUDED.id").
		Exec(r.ctx)
	return err
}
