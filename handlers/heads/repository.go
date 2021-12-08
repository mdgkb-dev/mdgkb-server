package heads

import (
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Head) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Heads, error) {
	items := make(models.Heads, 0)
	query := r.db.NewSelect().Model(&items).
		Relation("Human").
		Relation("Photo").
		Order("human.surname")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Head, error) {
	item := models.Head{}
	err := r.db.NewSelect().Model(&item).Where("heads.id = ?", id).
		Relation("Human").
		Relation("Photo").
		Relation("Regalias").
		Relation("Timetable.TimetableDays.Weekday").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Head{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Head) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
