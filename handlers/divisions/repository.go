package divisions

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.Division) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.DivisionsWithCount, err error) {
	item.Divisions = make(models.Divisions, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.Divisions).
		Relation("Entrance.Building").
		Relation("DivisionImages.FileInfo").
		Relation("Contact.Emails").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Contact.PostAddresses").
		Relation("Contact.Phones").
		Relation("Contact.Websites").
		Relation("MedicalProfilesDivisions.MedicalProfile").
		Relation("TreatDirection").
		Relation("Chief.Employee.Human")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Division, error) {
	item := models.Division{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("Entrance.Building").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Schedule.ScheduleItems").
		Relation("DivisionImages.FileInfo").
		Relation("DivisionVideos").
		Relation("DivisionPaidServices.PaidService").
		Relation("Comments", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("published_on DESC")
		}).
		Relation("Comments.User.Human").
		Relation("Timetable.TimetableDays.BreakPeriods").
		Relation("Timetable.TimetableDays.Weekday").
		// Relation("HospitalizationContactInfo.Emails").
		// Relation("HospitalizationContactInfo.TelephoneNumbers").
		Relation("Contact.Emails").
		Relation("Contact.PostAddresses").
		Relation("Contact.Phones").
		Relation("Contact.Websites").
		Relation("HospitalizationDoctor.Employee.Human").
		Relation("MedicalProfilesDivisions.MedicalProfile").
		Relation("TreatDirection").
		Relation("Chief.Employee.Human.Photo").
		Relation("NewsDivisions.News").
		Relation("DoctorsDivisions.Doctor.Employee.Human.PhotoMini").
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
		Where("divisions_view.id = ?", id).
		Scan(c)

	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Division{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Division) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) GetBySearch(c context.Context, search string) (models.Divisions, error) {
	items := make(models.Divisions, 0)

	err := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Column("divisions_view.id", "divisions_view.name", "divisions_view.slug").
		Where(r.helper.SQL.WhereLikeWithLowerTranslit("divisions_view.name", search)).
		Scan(c)
	return items, err
}
