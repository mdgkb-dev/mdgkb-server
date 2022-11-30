package pagesections

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.PageSection) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(params models.DocumentsParams) (items models.PageSections, err error) {
	query := r.db().NewSelect().Model(&items).Relation("DocumentTypeFields.ValueType")
	params.CreateJoin(query)
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.PageSection, error) {
	item := models.PageSection{}
	err := r.db().NewSelect().Model(&item).
		Relation("DocumentTypeFields.ValueType").
		Where("page_sections.id = ?", *id).
		Scan(context.Background())
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.PageSection{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PageSection) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.PageSection) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(item).
		Set("name = EXCLUDED.name").
		Set("description = EXCLUDED.description").
		Set("item_order = EXCLUDED.item_order").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PageSections) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("name = EXCLUDED.name").
		Set("description = EXCLUDED.description").
		Set("item_order = EXCLUDED.item_order").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.PageSection)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
