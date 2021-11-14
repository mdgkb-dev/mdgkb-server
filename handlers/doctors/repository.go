package doctors

import (
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Doctor) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(params *doctorsParams) (models.Doctors, error) {
	items := make(models.Doctors, 0)
	query := r.db.NewSelect().Model(&items).
		Relation("Human").
		Relation("Division").
		Relation("FileInfo").
		Order("human.surname")
	if params.Limit != 0 {
		query = query.Limit(params.Limit)
	}
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Doctor, error) {
	item := models.Doctor{}
	err := r.db.NewSelect().Model(&item).Where("doctor.id = ?", id).
		Relation("Human").
		Relation("FileInfo").
		Relation("Division").
		Relation("DoctorRegalias").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Educations.EducationCertification").
		Relation("Educations.EducationAccreditation").
		Relation("DoctorComments.Comment.User").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) getByDivisionID(id string) (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.db.NewSelect().
		Model(&items).
		Where("doctors.id = ?", id).
		Relation("Human").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Doctor{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Doctor) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) createComment(item *models.DoctorComment) error {
	_, err := r.db.NewInsert().Model(item.Comment).Exec(r.ctx)
	item.CommentId = item.Comment.ID
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) updateComment(item *models.DoctorComment) error {
	_, err := r.db.NewUpdate().Model(item.Comment).Where("id = ?", item.Comment.ID).Exec(r.ctx)
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) removeComment(id string) error {
	_, err := r.db.NewDelete().Model(&models.DoctorComment{}).Where("id = ?", id).Exec(r.ctx)
	return err
}
