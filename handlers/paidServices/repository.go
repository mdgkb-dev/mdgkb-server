package paidServices

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.PaidService) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.PaidServices, error) {
	items := make(models.PaidServices, 0)
	err := r.db().NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.PaidService, error) {
	item := models.PaidService{}
	err := r.db().NewSelect().
		Model(&item).
		Where("id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.PaidService{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PaidService) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getBySlug(slug *string) (*models.PaidService, error) {
	item := models.PaidService{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("PaidServiceItems").
		Where("slug = ?", *slug).
		Scan(r.ctx)
	return &item, err
}
