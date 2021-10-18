package educations

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.Educations) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.Education)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Educations) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("type = EXCLUDED.type").
		Set("institution = EXCLUDED.institution").
		Set("document = EXCLUDED.document").
		Set("qualification = EXCLUDED.qualification").
		Set("education_certification_id = EXCLUDED.education_certification_id").
		Set("education_accreditation_id = EXCLUDED.education_accreditation_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
