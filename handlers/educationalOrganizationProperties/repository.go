package educationalOrganizationProperties

import (

	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.EducationalOrganizationProperties, error) {
	items := make(models.EducationalOrganizationProperties, 0)
	err := r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.EducationalOrganizationProperty)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationalOrganizationProperties) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("name = EXCLUDED.name").
		Set("value = EXCLUDED.value").
		Model(&items).
		Exec(r.ctx)
	return err
}
