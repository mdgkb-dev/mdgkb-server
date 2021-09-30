package educationalOrganizationPages

import (

	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.EducationalOrganizationPages, error) {
	items := make(models.EducationalOrganizationPages, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("Page").Scan(r.ctx)
	return items, err
}


func (r *Repository) createMany(items models.EducationalOrganizationPages) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.EducationalOrganizationPage)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationalOrganizationPages) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("page_id = EXCLUDED.page_id").
		Exec(r.ctx)
	return err
}
