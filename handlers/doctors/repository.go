package doctors

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.Doctor) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.Doctors, err error) {
	query := r.db().NewSelect().Model(&items).
		Relation("DoctorsDivisions.Division.Floor").
		Relation("Position").
		Relation("MedicalProfile").
		Relation("Employee.Regalias").
		Relation("Employee.Human").
		Relation("Employee.Human.PhotoMini")
	r.queryFilter.HandleQuery(query)
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) getAllAdmin() (item models.DoctorsWithCount, err error) {
	item.Doctors = make(models.Doctors, 0)
	query := r.db().NewSelect().Model(&item.Doctors).
		Relation("DoctorsDivisions.Division.Floor").
		Relation("Position").
		Relation("MedicalProfile").
		Relation("Employee.Regalias").
		Relation("DoctorComments.Comment").
		Relation("Employee.Human")
	//Relation("Human.Photo").
	//Relation("Human.PhotoMini")
	// Join("JOIN positions on doctors_view.position_id = positions.id and positions.show = true")

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) getAllTimetables() (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(slug string) (*models.Doctor, error) {
	item := models.Doctor{}
	err := r.db().NewSelect().Model(&item).Where("doctors_view.slug = ?", slug).
		Relation("Employee.Human.Photo").
		Relation("Employee.Human.PhotoMini").
		Relation("DoctorsDivisions.Division.Timetable.TimetableDays.Weekday").
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
		//Relation("EducationalAcademic").
		Relation("Employee.TeachingActivities").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) getByDivisionID(id string) (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.db().NewSelect().
		Model(&items).
		Where("doctors_view.id = ?", id).
		Relation("Human").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Doctor{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Doctor) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) createComment(item *models.DoctorComment) error {
	_, err := r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) updateComment(item *models.DoctorComment) error {
	_, err := r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) removeComment(id string) error {
	_, err := r.db().NewDelete().Model(&models.DoctorComment{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Doctors) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("show = EXCLUDED.show").
		Set("has_appointment = EXCLUDED.has_appointment").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) search(search string) (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.db().NewSelect().
		Model(&items).
		Relation("Human").
		Where("lower(regexp_replace(human.name, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		WhereOr("lower(regexp_replace(human.surname, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		WhereOr("lower(regexp_replace(human.patronymic, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Doctor)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Doctor) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
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
		Exec(r.ctx)
	return err
}
