package doctors

import (
	"context"

	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.Doctor) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.Doctors, err error) {
	q := r.helper.DB.IDB(c).NewSelect().Model(&items).
		Relation("DoctorsDivisions.Division.Floor").
		Relation("Position").
		Relation("MedicalProfile").
		Relation("Employee.Regalias").
		Relation("Employee.Human").
		Relation("Employee.Human.PhotoMini")
	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	err = q.Scan(c)
	return items, err
}

func (r *Repository) GetAllTimetables(c context.Context) (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.helper.DB.IDB(c).NewSelect().Model(&items).
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context, slug string) (*models.Doctor, error) {
	item := models.Doctor{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).Where("doctors_view.slug = ?", slug).
		Relation("Employee.Human.Photo").
		Relation("Employee.Human.PhotoMini").
		// Relation("DoctorsDivisions.Division.Timetable.TimetableDays.Weekday").
		Relation("Employee.Regalias").
		Relation("Employee.Experiences").
		Relation("Position").
		Relation("DoctorPaidServices.PaidService").
		Relation("MedicalProfile").
		Relation("Employee.Certificates.Scan").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Relation("Employee.Educations").
		Relation("Employee.Certifications").
		Relation("Employee.Accreditations").
		Relation("DoctorComments.Comment", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("comment.mod_checked = true").Order("comment.published_on DESC")
		}).
		Relation("DoctorComments.Comment.User.Human").
		Relation("NewsDoctors.News").
		// Relation("EducationalAcademic").
		Relation("Employee.TeachingActivities").
		Scan(c)
	return &item, err
}

func (r *Repository) GetByDivisionID(c context.Context, id string) (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Where("doctors_view.id = ?", id).
		Relation("Human").
		Scan(c)
	return items, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Doctor{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Doctor) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) CreateComment(c context.Context, item *models.DoctorComment) error {
	_, err := r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) UpdateComment(c context.Context, item *models.DoctorComment) error {
	_, err := r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) RemoveComment(c context.Context, id string) error {
	_, err := r.helper.DB.IDB(c).NewDelete().Model(&models.DoctorComment{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.Doctors) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("show = EXCLUDED.show").
		Set("has_appointment = EXCLUDED.has_appointment").
		Model(&items).
		Exec(c)
	return err
}

func (r *Repository) DeleteMany(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.Doctor)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(c)
	return err
}

func (r *Repository) Upsert(c context.Context, item *models.Doctor) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("employee_id = EXCLUDED.employee_id").
		Set("position_id = EXCLUDED.position_id").
		Set("medical_profile_id = EXCLUDED.medical_profile_id").
		Set("timetable_id = EXCLUDED.timetable_id").
		Set("online_doctor_id = EXCLUDED.online_doctor_id").
		Set("mos_doctor_link = EXCLUDED.mos_doctor_link").
		Set("show = EXCLUDED.show").
		Set("has_appointment = EXCLUDED.has_appointment").
		Model(item).
		Exec(c)
	return err
}
