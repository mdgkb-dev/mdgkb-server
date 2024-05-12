package users

import (
	"context"
	"mdgkb/mdgkb-server/models"

	//_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

func (r *Repository) GetAll(c context.Context) (item models.UsersWithCount, err error) {
	item.Users = make(models.Users, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&item.Users).
		Relation("Human").
		Relation("Role")

	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id string) (*models.User, error) {
	item := models.User{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("Human.Photo").
		Relation("Human.Contact.Address").
		Relation("Questions.User.Human").
		Relation("Comments").
		Relation("Comments.User.Human").
		Relation("Questions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("questions.question_date desc")
		}).
		Relation("DonorRulesUsers.DonorRule.Image").
		Relation("DonorRulesUsers.DonorRule.DonorRulesUsers").
		Relation("DoctorsUsers.Doctor").
		Relation("Children.Human").
		Relation("DailyMenuOrders.DailyMenuOrderItems.DailyMenuItem").
		Relation("DailyMenuOrders", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("daily_menu_orders_view.created_at desc")
		}).
		Relation("DailyMenuOrders.FormValue.User.Human").
		Relation("DailyMenuOrders.FormValue.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("DailyMenuOrders.FormValue.Fields.File").
		Relation("DailyMenuOrders.FormValue.FormValueFiles.File").
		Relation("DailyMenuOrders.FormValue.Fields.ValueType").
		Relation("DailyMenuOrders.FormValue.FieldValues.File").
		Relation("DailyMenuOrders.FormValue.FieldValues.Field.ValueType").
		Relation("DailyMenuOrders.FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("Role").
		Relation("ResidencyApplications.ResidencyCourse.ResidencyCoursesSpecializations.Specialization").
		Relation("ResidencyApplications.FormValue.Fields.File").
		Relation("ResidencyApplications.FormValue.FormValueFiles.File").
		Relation("ResidencyApplications.FormValue.Fields.ValueType").
		Relation("ResidencyApplications.FormValue.FieldValues.File").
		Relation("ResidencyApplications.FormValue.FieldValues.Field.ValueType").
		Relation("ResidencyApplications.FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("ResidencyApplications.FormValue.User").
		Relation("VacancyResponses.Vacancy").
		Relation("VacancyResponses.FormValue.Fields.File").
		Relation("VacancyResponses.FormValue.FormValueFiles.File").
		Relation("VacancyResponses.FormValue.Fields.ValueType").
		Relation("VacancyResponses.FormValue.FieldValues.File").
		Relation("VacancyResponses.FormValue.FieldValues.Field.ValueType").
		Relation("VacancyResponses.FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("VacancyResponses.FormValue.User").
		//Relation("FormValues.User").
		//Relation("FormValues.FormValueFiles.File").
		//Relation("FormValues.FieldValues.Field").
		//Relation("FormValues.FieldValues.File").
		//Relation("FormValues.Fields.File").
		//Relation("FormValues.FormStatus.FormStatusToFormStatuses.ChildFormStatus.Icon").
		//Relation("FormValues.DpoApplication.NmoCourse"). // TODO: исправить столбец в базе, когда будут заявки
		// Relation("FormValues.PostgraduateApplication.PostgraduateCourse", func(query *bun.SelectQuery) *bun.SelectQuery {
		// 	return query.ExcludeColumn("questions_file_id")
		// }).
		// Relation("FormValues.PostgraduateApplication.PostgraduateCourse.PostgraduateCoursesSpecializations.Specialization").
		// Relation("FormValues.ResidencyApplication.ResidencyCourse.ResidencyCoursesSpecializations.Specialization").
		// Relation("FormValues.CandidateApplication.CandidateExam").
		// Relation("FormValues.VacancyResponse.Vacancy").
		// Relation("FormValues.VisitsApplication.Division").
		// Relation("FormValues.DailyMenuOrder.DailyMenuOrderItems.DailyMenuItem").
		Where("users_view.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) GetByEmail(c context.Context, id string) (*models.User, error) {
	item := models.User{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Human.Photo").
		Relation("Human.Contact.Address").
		//Relation("Questions").
		Relation("DonorRulesUsers.DonorRule.Image").
		Relation("DonorRulesUsers.DonorRule.DonorRulesUsers").
		Relation("Children.Human").
		Relation("Role").
		Relation("DoctorsUsers").
		Where("users_view.email = ? AND users_view.is_active = true", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Create(c context.Context, user *models.User) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(user).Exec(c)
	return err
}

func (r *Repository) EmailExists(c context.Context, email string) (bool, error) {
	exists, err := r.helper.DB.IDB(c).NewSelect().Model((*models.User)(nil)).Where("users_view.email = ? and is_active = true", email).Exists(c)
	return exists, err
}

func (r *Repository) Update(c context.Context, item *models.User) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).
		OmitZero().
		ExcludeColumn("password", "is_active"). // all columns except col1
		Where("id = ?", item.ID).
		Exec(c)
	return err
}

func (r *Repository) Upsert(c context.Context, item *models.User) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (email) do update").Model(item).
		Set("role_id = EXCLUDED.role_id").
		Set("password = EXCLUDED.password").
		Set("is_active = EXCLUDED.is_active").
		Exec(c)
	return err
}

func (r *Repository) UpsertEmail(c context.Context, item *models.User) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (email) DO UPDATE").
		Set("phone = EXCLUDED.phone").
		Model(item).
		Exec(c)
	return err
}

func (r *Repository) AddToUser(c context.Context, values map[string]interface{}, table string) error {
	_, err := r.helper.DB.IDB(c).NewInsert().Model(&values).TableExpr(table).Exec(c)
	return err
}

func (r *Repository) RemoveFromUser(c context.Context, values map[string]interface{}, table string) error {
	q := r.helper.DB.IDB(c).NewDelete().Table(table)
	for key, value := range values {
		q = q.Where("? = ?", bun.Ident(key), value)
	}
	_, err := q.Exec(c)
	return err
}

func (r *Repository) DropUUID(c context.Context, item *models.User) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().
		Model(item).
		Set("uuid = uuid_generate_v4()").
		Where("id = ?", item.ID).
		Exec(c)
	return err
}

func (r *Repository) UpdatePassword(c context.Context, item *models.User) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().
		Model(item).
		Set("password = ?", item.Password).
		Set("is_active = true").
		Where("id = ?", item.ID).
		Exec(c)
	return err
}
func (r *Repository) GetByUserAccountID(c context.Context, id string) (*models.User, error) {
	item := models.User{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Role").
		Where("?TableAlias.user_account_id = ?", id).
		Scan(c)
	return &item, err
}
