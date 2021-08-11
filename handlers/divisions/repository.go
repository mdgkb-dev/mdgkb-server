package divisions

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.Division) error
	getAll(*gin.Context) ([]models.Division, error)
	get(*gin.Context, string) (models.Division, error)
	updateStatus(*gin.Context, *models.Division) error
	delete(*gin.Context, string) error
	update(*gin.Context, *models.Division) error
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, item *models.Division) (err error) {
	_, err = r.db.NewInsert().Model(item.Timetable).On("conflict (id) do update").
		Set("description = ?", item.Timetable.Description).Exec(ctx)

	item.Timetable.SetIdForChildren()

	item.TimetableId = item.Timetable.ID

	_, err = r.db.NewInsert().Model(item).Exec(ctx)

	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&item.Timetable.TimetableDays).
		Set("is_weekend = EXCLUDED.is_weekend").
		Set("timetable_id = EXCLUDED.timetable_id").
		Set("weekday_id = EXCLUDED.weekday_id").
		Set("start_time = EXCLUDED.start_time").
		Set("end_time = EXCLUDED.end_time").
		Set("break_exist = EXCLUDED.break_exist").
		Set("break_start_time = EXCLUDED.break_start_time").
		Set("break_end_time = EXCLUDED.break_end_time").
		Where("timetable_day.id = EXCLUDED.id").
		Exec(ctx)

	return err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.Division, err error) {
	err = r.db.NewSelect().Model(&items).Order("name").Scan(ctx)
	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.Division, err error) {
	err = r.db.NewSelect().
		Model(&item).
		Relation("Timetable.TimetableDays.Weekday").
		Where("division.id = ?", id).Scan(ctx)
	err = r.db.NewSelect().Model(&item.Doctors).Where("division_id = ?", id).
		Relation("FileInfo").
		Relation("Human").
		Scan(ctx)
	return item, err
}

func (r *Repository) updateStatus(ctx *gin.Context, item *models.Division) (err error) {
	_, err = r.db.NewUpdate().Model(item).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Division{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *Repository) update(ctx *gin.Context, item *models.Division) (err error) {
	_, err = r.db.NewInsert().Model(item.Timetable).On("conflict (id) do update").
		Set("description = ?", item.Timetable.Description).Exec(ctx)
	fmt.Println(err)
	item.Timetable.SetIdForChildren()

	item.TimetableId = item.Timetable.ID
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(ctx)
	fmt.Println(err)
	if len(item.Timetable.TimetableDaysForDelete) > 0 {
		_, err = r.db.NewDelete().Model((*models.TimetableDay)(nil)).Where("id IN (?)", bun.In(item.Timetable.TimetableDaysForDelete)).Exec(ctx)
	}
	fmt.Println(err)
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&item.Timetable.TimetableDays).
		Set("is_weekend = EXCLUDED.is_weekend").
		Set("timetable_id = EXCLUDED.timetable_id").
		Set("weekday_id = EXCLUDED.weekday_id").
		Set("start_time = EXCLUDED.start_time").
		Set("end_time = EXCLUDED.end_time").
		Set("break_exist = EXCLUDED.break_exist").
		Set("break_start_time = EXCLUDED.break_start_time").
		Set("break_end_time = EXCLUDED.break_end_time").
		Where("timetable_day.id = EXCLUDED.id").
		Exec(ctx)
	fmt.Println(err)
	return err
}
