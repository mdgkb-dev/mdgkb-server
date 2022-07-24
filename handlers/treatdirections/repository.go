package treatdirections

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
	//_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.TreatDirection) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.TreatDirections, error) {
	items := make(models.TreatDirections, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("Division").
		Order("treat_directions.name")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.TreatDirection, error) {
	item := models.TreatDirection{}
	err := r.db().NewSelect().Model(&item).
		Relation("Division").
		Where("treat_directions.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.TreatDirection{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.TreatDirection) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}