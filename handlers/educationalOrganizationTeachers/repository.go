package educationalOrganizationTeachers

import (

	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.EducationalOrganizationTeachers, error) {
	items := make(models.EducationalOrganizationTeachers, 0)
	err := r.db.NewSelect().Model(&items).Relation("Doctor.Human").Scan(r.ctx)
	return items, err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.EducationalOrganizationTeacher)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationalOrganizationTeachers) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("doctor_id = EXCLUDED.doctor_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
