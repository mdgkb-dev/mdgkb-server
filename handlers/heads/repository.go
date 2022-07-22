package heads

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
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
		Relation("Human").
		Relation("Photo").
		Relation("Departments.Division").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites")
	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Head, error) {
	item := models.Head{}
	err := r.DB().NewSelect().Model(&item).Where("heads.id = ?", id).
		Relation("Human").
		Relation("Photo").
		Relation("Regalias").
		Relation("Departments.Division").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
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
