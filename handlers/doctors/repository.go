package doctors

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
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

func (r *Repository) create(item *models.Doctor) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAllMain() (items models.Doctors, err error) {
	err = r.db.NewSelect().Model(&items).
		Relation("Human").
		Relation("Division.Floor").
		Relation("FileInfo").
		Relation("PhotoMini").
		Relation("Position").
		Relation("MedicalProfile").
		Relation("Regalias").
		Relation("DoctorComments.Comment").
		// Join("JOIN positions on doctors_view.position_id = positions.id and positions.show = true").
		Order("doctors_view.regalias_count DESC", "doctors_view.comments_count DESC").
		Where("doctors_view.file_info_id is not null").
		Where("doctors_view.mos_doctor_link is not null and doctors_view.mos_doctor_link != '' ").
		Limit(20).Scan(r.ctx)
	return items, err
}

func (r *Repository) getAll() (items models.Doctors, err error) {
	query := r.db.NewSelect().Model(&items).
		Relation("Human").
		Relation("Division.Floor").
		Relation("FileInfo").
		Relation("PhotoMini").
		Relation("Position").
		Relation("MedicalProfile").
		Relation("Regalias").
		Relation("DoctorComments.Comment")
	// Join("JOIN positions on doctors_view.position_id = positions.id and positions.show = true")

	r.queryFilter.HandleQuery(query)
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) getAllAdmin() (items models.DoctorsWithCount, err error) {
	query := r.db.NewSelect().Model(&items.Doctors).
		Relation("Human").
		Relation("Division.Floor").
		Relation("FileInfo").
		Relation("PhotoMini").
		Relation("Position").
		Relation("MedicalProfile").
		Relation("Regalias").
		Relation("DoctorComments.Comment")
	// Join("JOIN positions on doctors_view.position_id = positions.id and positions.show = true")

	r.queryFilter.HandleQuery(query)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) getAllTimetables() (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(slug string) (*models.Doctor, error) {
	item := models.Doctor{}
	err := r.db.NewSelect().Model(&item).Where("doctors_view.slug = ?", slug).
		Relation("Human").
		Relation("FileInfo").
		Relation("PhotoMini").
		Relation("Division.Timetable.TimetableDays.Weekday").
		Relation("Regalias").
		Relation("Experiences").
		Relation("Position").
		Relation("DoctorPaidServices.PaidService").
		Relation("MedicalProfile").
		Relation("Certificates.Scan").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Relation("Educations.EducationCertification").
		Relation("Educations.EducationAccreditation").
		Relation("DoctorComments.Comment.User.Human").
		Relation("NewsDoctors.News").
		Relation("EducationalOrganizationAcademic").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) getByDivisionID(id string) (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.db.NewSelect().
		Model(&items).
		Where("doctors_view.id = ?", id).
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

func (r *Repository) upsertMany(items models.Doctors) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("show = EXCLUDED.show").
		Set("division_id = EXCLUDED.division_id").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) search(search string) (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.db.NewSelect().
		Model(&items).
		Relation("Human").
		Where("lower(regexp_replace(human.name, '[^??-????-??a-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		WhereOr("lower(regexp_replace(human.surname, '[^??-????-??a-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		WhereOr("lower(regexp_replace(human.patronymic, '[^??-????-??a-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		Scan(r.ctx)
	return items, err
}
