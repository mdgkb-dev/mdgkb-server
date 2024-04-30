package heads

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.Head) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Heads, error) {
	items := make(models.Heads, 0)
	query := r.DB().NewSelect().Model(&items).
		Relation("Employee.Human.Photo").
		Relation("Employee.Human.PhotoMini").
		Relation("Departments.Division").
		Relation("Timetable.TimetableDays.Weekday")
		// Relation("ContactInfo").
		// Relation("ContactInfo.Emails").
		// Relation("ContactInfo.PostAddresses").
		// Relation("ContactInfo.TelephoneNumbers").
		// Relation("ContactInfo.Websites")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Head, error) {
	item := models.Head{}
	err := r.DB().NewSelect().Model(&item).Where("?TableAlias.id = ?", id).
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
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.Head{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Head) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) updateAll(items models.Heads) (err error) {
	_, err = r.DB().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("item_order = EXCLUDED.item_order").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Head) (err error) {
	_, err = r.DB().NewInsert().On("conflict (id) do update").
		Model(item).
		Set("employee_id = EXCLUDED.employee_id").
		Set("position = EXCLUDED.position").
		Set("timetable_id = EXCLUDED.timetable_id").
		Set("is_main = EXCLUDED.is_main").
		Set("contact_info_id = EXCLUDED.contact_info_id").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.DB().NewDelete().
		Model((*models.Head)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
