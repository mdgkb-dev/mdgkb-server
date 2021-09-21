package divisions

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.Division) error
	getAll(*gin.Context) ([]models.Division, error)
	get(*gin.Context, string) (models.Division, error)
	delete(*gin.Context, string) error
	update(*gin.Context, *models.Division) error
	createComment(*gin.Context, *models.DivisionComment) error
	updateComment(*gin.Context, *models.DivisionComment) error
	removeComment(*gin.Context, string) error
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
	if err != nil {
		return err
	}

	_, err = r.db.NewInsert().Model(item.Schedule).On("conflict (id) do update").
		Set("name = EXCLUDED.name").Exec(ctx)
	item.Schedule.SetIdForChildren()
	item.ScheduleId = item.Schedule.ID
	if err != nil {
		return err
	}

	_, err = r.db.NewInsert().Model(item).Exec(ctx)
	if err != nil {
		return err
	}

	if len(item.Timetable.TimetableDays) > 0 {
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
		if err != nil {
			return err
		}
	}

	if len(item.Schedule.ScheduleItems) > 0 {
		_, err = r.db.NewInsert().On("conflict (id) do update").
			Model(&item.Schedule.ScheduleItems).
			Set("name = EXCLUDED.name").
			Set("start_time = EXCLUDED.start_time").
			Set("end_time = EXCLUDED.end_time").
			Where("schedule_item.id = EXCLUDED.id").
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	for _, divisionImage := range item.DivisionImages {
		_, err = r.db.NewInsert().Model(divisionImage.FileInfo).Exec(ctx)
		divisionImage.FileInfoId = divisionImage.FileInfo.ID
		divisionImage.DivisionId = item.ID
		_, err = r.db.NewInsert().Model(divisionImage).Exec(ctx)
	}

	return err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.Division, err error) {
	err = r.db.NewSelect().Model(&items).
		Relation("Entrance.Building").
		// Relation("DivisionImages.FileInfo").
		Order("name").
		Scan(ctx)
	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.Division, err error) {
	err = r.db.NewSelect().
		Model(&item).
		Relation("Entrance.Building").
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Schedule.ScheduleItems").
		Relation("DivisionImages.FileInfo").
		Relation("DivisionComments.Comment.User").
		Where("division.id = ?", id).Scan(ctx)

	err = r.db.NewSelect().Model(&item.Doctors).Where("division_id = ?", id).
		Relation("FileInfo").
		Relation("Human").
		Scan(ctx)
	return item, err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Division{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *Repository) update(ctx *gin.Context, item *models.Division) (err error) {
	_, err = r.db.NewInsert().Model(item.Timetable).On("conflict (id) do update").
		Set("description = ?", item.Timetable.Description).Exec(ctx)
	item.Timetable.SetIdForChildren()
	item.TimetableId = item.Timetable.ID
	if err != nil {
		return err
	}

	_, err = r.db.NewInsert().Model(item.Schedule).On("conflict (id) do update").
		Set("description = EXCLUDED.description").Exec(ctx)
	item.Schedule.SetIdForChildren()
	item.ScheduleId = item.Schedule.ID
	if err != nil {
		return err
	}

	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(ctx)
	if len(item.Timetable.TimetableDaysForDelete) > 0 {
		_, err = r.db.NewDelete().Model((*models.TimetableDay)(nil)).Where("id IN (?)", bun.In(item.Timetable.TimetableDaysForDelete)).Exec(ctx)
	}
	if err != nil {
		return err
	}

	if len(item.Schedule.ScheduleItemsForDelete) > 0 {
		_, err = r.db.NewDelete().Model((*models.ScheduleItem)(nil)).Where("id IN (?)", bun.In(item.Schedule.ScheduleItemsForDelete)).Exec(ctx)
	}
	if err != nil {
		return err
	}

	if len(item.Timetable.TimetableDays) > 0 {
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
	}

	if len(item.Schedule.ScheduleItems) > 0 {
		_, err = r.db.NewInsert().On("conflict (id) do update").
			Model(&item.Schedule.ScheduleItems).
			Set("name = EXCLUDED.name").
			Set("start_time = EXCLUDED.start_time").
			Set("end_time = EXCLUDED.end_time").
			Where("schedule_item.id = EXCLUDED.id").
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	if len(item.DivisionImagesForDelete) > 0 {
		_, err = r.db.NewDelete().Model((*models.NewsImage)(nil)).Where("id IN (?)", bun.In(item.DivisionImagesForDelete)).Exec(ctx)
	}
	if len(item.DivisionImages) == 0 {
		return err
	}
	var fileInfos []models.FileInfo
	for _, newsImage := range item.DivisionImages {
		fileInfos = append(fileInfos, *newsImage.FileInfo)
	}

	_, err = r.db.NewInsert().Model(&fileInfos).
		On("CONFLICT (id) DO UPDATE").
		Set("original_name = EXCLUDED.original_name").
		Set("file_system_path = EXCLUDED.file_system_path").
		Exec(ctx)

	for i, newsImage := range item.DivisionImages {
		newsImage.FileInfoId = fileInfos[i].ID
		newsImage.DivisionId = item.ID
	}

	_, err = r.db.NewInsert().Model(&item.DivisionImages).On("CONFLICT (id) DO UPDATE").
		Set("description = EXCLUDED.description").
		Set("file_info_id = EXCLUDED.file_info_id").
		Exec(ctx)

	return err
}

func (r *Repository) createComment(ctx *gin.Context, item *models.DivisionComment) error {
	_, err := r.db.NewInsert().Model(item.Comment).Exec(ctx)
	item.CommentId = item.Comment.ID
_, err = r.db.NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *Repository) updateComment(ctx *gin.Context, item *models.DivisionComment) error {
	_, err := r.db.NewUpdate().Model(item.Comment).Where("id = ?", item.Comment.ID).Exec(ctx)
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(ctx)
	return err
}

func (r *Repository) removeComment(ctx *gin.Context, id string) error {
	_, err := r.db.NewDelete().Model(&models.DivisionComment{}).Where("id = ?", id).Exec(ctx)
	return err
}		
