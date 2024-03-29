package dietsgroups

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.DietGroup) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.DietsGroups, error) {
	items := make(models.DietsGroups, 0)
	query := r.db().NewSelect().Model(&items).Order("diets_groups.diet_group_order").
		Relation("Diets", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("diets.diet_order")
		}).
		Relation("Diets.DietAges", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("diet_ages.diet_age_order")
		}).
		Relation("Diets.DietAges.Timetable.TimetableDays.ScheduleItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("schedule_item.schedule_item_order")
		}).
		Relation("Diets.DietAges.Timetable.TimetableDays.ScheduleItems.Dishes")
	//Relation("Diets.MotherDiet.DietAges.Timetable.TimetableDays.ScheduleItems.Dishes")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.DietGroup, error) {
	item := models.DietGroup{}
	err := r.db().NewSelect().Model(&item).
		Where("diets.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.DietGroup{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DietGroup) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
