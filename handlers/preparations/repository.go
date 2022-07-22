package preparations

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Preparation) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Preparations, error) {
	items := make(models.Preparations, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("PreparationRulesGroups.PreparationRules").
		Relation("PreparationsToTags").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Preparation, error) {
	item := models.Preparation{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("PreparationRulesGroups.PreparationRules").
		Where("preparations.id = ?", id).
		Scan(r.ctx)

	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Preparation{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Preparation) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Preparations) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Preparation)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) getTags() (models.PreparationsTags, error) {
	items := make(models.PreparationsTags, 0)
	err := r.db().NewSelect().Model(&items).
		Scan(r.ctx)
	return items, err
}
