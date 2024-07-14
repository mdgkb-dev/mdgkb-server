package heads

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.Head) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (models.Heads, error) {
	items := make(models.Heads, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&items).
		Relation("Employee.Human.Photo").
		Relation("Employee.Human.PhotoMini").
		Relation("Departments.Division").
		Relation("Timetable.TimetableDays.Weekday")
		// Relation("ContactInfo").
		// Relation("ContactInfo.Emails").
		// Relation("ContactInfo.PostAddresses").
		// Relation("ContactInfo.TelephoneNumbers").
		// Relation("ContactInfo.Websites")
	err := query.Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Head, error) {
	item := models.Head{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).Where("?TableAlias.id = ?", id).
		Relation("Employee.Human.Photo").
		Relation("Employee.Human.PhotoMini").
		Relation("Employee.Regalias").
		Relation("Departments.Division").
		Relation("Timetable.TimetableDays.Weekday").
		// Relation("ContactInfo").
		// Relation("ContactInfo.Emails").
		// Relation("ContactInfo.PostAddresses").
		// Relation("ContactInfo.TelephoneNumbers").
		// Relation("ContactInfo.Websites").
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Head{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Head) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) UpdateAll(c context.Context, items models.Heads) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("item_order = EXCLUDED.item_order").
		Exec(c)
	return err
}

func (r *Repository) Upsert(c context.Context, item *models.Head) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(item).
		Set("employee_id = EXCLUDED.employee_id").
		Set("position = EXCLUDED.position").
		Set("timetable_id = EXCLUDED.timetable_id").
		Set("is_main = EXCLUDED.is_main").
		Set("contact_info_id = EXCLUDED.contact_info_id").
		Exec(c)
	return err
}

func (r *Repository) DeleteMany(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.Head)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(c)
	return err
}
