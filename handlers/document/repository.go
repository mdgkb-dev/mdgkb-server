package document

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Document) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.Documents, err error) {
	err = r.db.NewSelect().Model(&items).Relation("DocumentFields").Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Document, error) {
	item := models.Document{}
	err := r.db.NewSelect().Model(&item).
		Relation("DocumentFields").
		Where("documents.id = ?", *id).
		Scan(context.Background())
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Document{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Document) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
