package doctors

import (
	"github.com/gin-gonic/gin"
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

func (r *Repository) getAll(params *doctorsParams) (items models.DoctorsWithCount, err error) {
	query := r.db.NewSelect().Model(&items.Doctors).
		Relation("Human").
		Relation("Division").
		Relation("FileInfo").
		Relation("Position").
		Relation("MedicalProfile").
		Order("human.surname")

	if r.queryFilter != nil && r.queryFilter.Pagination != nil {
		r.helper.HTTP.CreatePaginationQuery(query, r.queryFilter.Pagination)
	}
	//r.helper.HTTP.CreateFilter(query, r.queryFilter.FilterModels)

	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) get(slug string) (*models.Doctor, error) {
	item := models.Doctor{}
	err := r.db.NewSelect().Model(&item).Where("doctors_view.slug = ?", slug).
		Relation("Human").
		Relation("FileInfo").
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
		Relation("DoctorComments.Comment.User").
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

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.HTTP.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) search(search string) (models.Doctors, error) {
	items := make(models.Doctors, 0)
	err := r.db.NewSelect().
		Model(&items).
		Relation("Human").
		Where("lower(regexp_replace(human.name, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		WhereOr("lower(regexp_replace(human.surname, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		WhereOr("lower(regexp_replace(human.patronymic, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		Scan(r.ctx)
	return items, err
}
