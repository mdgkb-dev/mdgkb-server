package users

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.Users, error) {
	items := make(models.Users, 0)
	err := r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.User, error) {
	item := models.User{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("Human.Photo").
		//Relation("Questions").
		Relation("DonorRulesUsers.DonorRule.Image").
		Relation("DonorRulesUsers.DonorRule.DonorRulesUsers").
		Relation("DoctorsUsers.Doctor").
		Relation("Children.Human").
		Relation("FormValues.User").
		Relation("FormValues.FormStatus.FormStatusToFormStatuses.ChildFormStatus.Icon").
		Relation("FormValues.DpoApplication.DpoCourse").
		Relation("FormValues.PostgraduateApplication.PostgraduateCourse", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.ExcludeColumn("questions_file_id")
		}).
		Relation("FormValues.PostgraduateApplication.PostgraduateCourse.PostgraduateCoursesSpecializations.Specialization").
		Relation("FormValues.ResidencyApplication.ResidencyCourse.ResidencyCoursesSpecializations.Specialization").
		Relation("FormValues.CandidateApplication.CandidateExam").
		Where("users.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) getByEmail(id string) (*models.User, error) {
	item := models.User{}
	err := r.db.NewSelect().Model(&item).
		Relation("Human.Photo").
		//Relation("Questions").
		Relation("DonorRulesUsers.DonorRule.Image").
		Relation("DonorRulesUsers.DonorRule.DonorRulesUsers").
		Relation("Children.Human").
		Relation("Role").
		Relation("DoctorsUsers").
		Where("users.email = ? AND users.is_active = true", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(user *models.User) (err error) {
	_, err = r.db.NewInsert().Model(user).Exec(r.ctx)
	return err
}

func (r *Repository) emailExists(email string) (bool, error) {
	exists, err := r.db.NewSelect().Model((*models.User)(nil)).Where("users.email = ? and is_active = true", email).Exists(r.ctx)
	return exists, err
}

func (r *Repository) update(item *models.User) (err error) {
	_, err = r.db.NewUpdate().Model(item).OmitZero().
		Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.User) (err error) {
	_, err = r.db.NewInsert().On("conflict (email) do update").Model(item).
		Set("role_id = EXCLUDED.role_id").
		Set("password = EXCLUDED.password").
		Set("is_active = EXCLUDED.is_active").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertEmail(item *models.User) (err error) {
	_, err = r.db.NewInsert().On("conflict (email) do update").
		Model(item).
		Exec(r.ctx)
	return err
}

func (r *Repository) addToUser(values map[string]interface{}, table string) error {
	_, err := r.db.NewInsert().Model(&values).TableExpr(table).Exec(r.ctx)
	return err
}

func (r *Repository) removeFromUser(values map[string]interface{}, table string) error {
	q := r.db.NewDelete().Table(table)
	for key, value := range values {
		q = q.Where(fmt.Sprintf("%s = ?", key), value)
	}
	_, err := q.Exec(r.ctx)
	return err
}

func (r *Repository) dropUUID(item *models.User) (err error) {
	_, err = r.db.NewUpdate().
		Model(item).
		Set("uuid = uuid_generate_v4()").
		Where("id = ?", item.ID).
		Exec(r.ctx)
	return err
}

func (r *Repository) updatePassword(item *models.User) (err error) {
	_, err = r.db.NewUpdate().
		Model(item).
		Set("password = ?", item.Password).
		Set("is_active = true").
		Where("id = ?", item.ID).
		Exec(r.ctx)
	return err
}
