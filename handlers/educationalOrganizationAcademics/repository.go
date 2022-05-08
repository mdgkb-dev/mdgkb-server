package educationalOrganizationAcademics

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.EducationalOrganizationAcademics, error) {
	items := make(models.EducationalOrganizationAcademics, 0)
	// TODO: panic from relation FileInfo&PhotoMini
	query := r.db.NewSelect().Model(&items).
		Relation("Doctor.Human.Photo").
		//Relation("Doctor.FileInfo").
		//Relation("Doctor.PhotoMini").
		Relation("Doctor.Position").
		Relation("Doctor.MedicalProfile").
		Relation("Doctor.Regalias").
		Relation("Doctor.DoctorComments.Comment")
	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.EducationalOrganizationAcademic)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationalOrganizationAcademics) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("doctor_id = EXCLUDED.doctor_id").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.EducationalOrganizationAcademic) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("doctor_id = EXCLUDED.doctor_id").
		Model(item).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteByDoctorID(id uuid.NullUUID) (err error) {
	_, err = r.db.NewDelete().Model(&models.EducationalOrganizationAcademic{}).Where("doctor_id = ?", id).Exec(r.ctx)
	return err
}
