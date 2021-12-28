package partners

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.Partners, error) {
	items := make(models.Partners, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("Image").
		Relation("PartnerType").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Partner, error) {
	item := models.Partner{}
	err := r.db.NewSelect().Model(&item).
		Relation("Image").
		Relation("PartnerType").
		Where("partners.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.Partner) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Partner{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Partner) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
