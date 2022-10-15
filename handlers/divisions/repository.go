package divisions

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Division) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (item models.DivisionsWithCount, err error) {
	item.Divisions = make(models.Divisions, 0)
	query := r.db().NewSelect().Model(&item.Divisions).
		Relation("Entrance.Building").
		Relation("DivisionImages.FileInfo").
		Relation("ContactInfo.Emails").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Relation("MedicalProfilesDivisions.MedicalProfile").
		Relation("TreatDirection").
		Relation("Chief.Human")

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get() (*models.Division, error) {
	item := models.Division{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("Entrance.Building").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Schedule.ScheduleItems").
		Relation("DivisionImages.FileInfo").
		Relation("DivisionPaidServices.PaidService").
		Relation("DivisionComments.Comment.User.Human", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("comment.published_on DESC")
		}).
		Relation("Timetable.TimetableDays.BreakPeriods").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("HospitalizationContactInfo.Emails").
		Relation("HospitalizationContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Relation("HospitalizationDoctor.Human").
		Relation("MedicalProfilesDivisions.MedicalProfile").
		Relation("TreatDirection").
		Relation("Chief.Human.Photo").
		Relation("NewsDivisions.News").
		Relation("DoctorsDivisions.Doctor.Human.PhotoMini").
		Relation("DoctorsDivisions.Doctor.Position").
		Relation("DoctorsDivisions.Doctor").
		Relation("DoctorsDivisions.Doctor.MedicalProfile").
		Relation("Vacancies").
		Relation("VisitingRulesGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("visiting_rules_groups.visiting_rule_group_order")
		}).
		Relation("VisitingRulesGroups.VisitingRules", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("visiting_rules.rule_order")
		}).
		Where("divisions_view.? = ?", bun.Safe(r.queryFilter.Col), r.queryFilter.Value).
		Scan(r.ctx)

	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Division{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Division) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) createComment(item *models.DivisionComment) error {
	_, err := r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) updateComment(item *models.DivisionComment) error {
	_, err := r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) removeComment(id string) error {
	_, err := r.db().NewDelete().Model(&models.DivisionComment{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) getBySearch(search string) (models.Divisions, error) {
	items := make(models.Divisions, 0)

	err := r.db().NewSelect().
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
