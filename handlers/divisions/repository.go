package divisions

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Division) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(onlyShowed bool) (items models.Divisions, err error) {
	query := r.db.NewSelect().Model(&items).
		Relation("Entrance.Building").
		Relation("DivisionImages.FileInfo").
		Relation("ContactInfo.Emails").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Relation("MedicalProfilesDivisions.MedicalProfile")

	if onlyShowed {
		query = query.Where("divisions_view.show = true")
	}
	r.queryFilter.HandleQuery(query)
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(slug string, onlyShowed bool) (*models.Division, error) {
	item := models.Division{}
	q := r.db.NewSelect().
		Model(&item).
		Relation("Entrance.Building").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Schedule.ScheduleItems").
		Relation("DivisionImages.FileInfo").
		Relation("DivisionPaidServices.PaidService").
		Relation("DivisionComments.Comment.User").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("HospitalizationContactInfo.Emails").
		Relation("HospitalizationContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Relation("HospitalizationDoctor.Human").
		Relation("MedicalProfilesDivisions.MedicalProfile")
	//if onlyShowed {
	q = q.Relation("Doctors", func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.
			Join("JOIN positions on doctors_view.position_id = positions.id and positions.show = true").
			Order("positions.item_order")
	})
	//}

	err := q.Relation("Doctors.FileInfo").
		Relation("Doctors.Human").
		Relation("Doctors.Position").
		Relation("Doctors.Division").
		Relation("Doctors.MedicalProfile").
		Relation("Vacancies").
		Relation("VisitingRules").
		Where("divisions_view.id = ?", slug).
		Scan(r.ctx)

	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Division{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Division) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) createComment(item *models.DivisionComment) error {
	_, err := r.db.NewInsert().Model(item.Comment).Exec(r.ctx)
	item.CommentId = item.Comment.ID
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) updateComment(item *models.DivisionComment) error {
	_, err := r.db.NewUpdate().Model(item.Comment).Where("id = ?", item.Comment.ID).Exec(r.ctx)
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) removeComment(id string) error {
	_, err := r.db.NewDelete().Model(&models.DivisionComment{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) getBySearch(search string) (models.Divisions, error) {
	items := make(models.Divisions, 0)

	err := r.db.NewSelect().
		Model(&items).
		Column("divisions_view.id", "divisions_view.name", "divisions_view.slug").
		Where(r.helper.SQL.WhereLikeWithLowerTranslit("divisions_view.name", search)).
		Scan(r.ctx)
	return items, err
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}
