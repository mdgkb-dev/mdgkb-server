package doctors

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.Doctor) error
	getAll(*gin.Context) ([]models.Doctor, error)
	get(*gin.Context, string) (models.Doctor, error)
	getByDivisionId(*gin.Context, string) ([]models.Doctor, error)
	updateStatus(*gin.Context, *models.Doctor) error
	delete(*gin.Context, string) error
	update(*gin.Context, *models.Doctor) error
	createComment(*gin.Context, *models.DoctorComment) error
	updateComment(*gin.Context, *models.DoctorComment) error
	removeComment(*gin.Context, string) error
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, item *models.Doctor) (err error) {
	_, err = r.db.NewInsert().Model(item.FileInfo).Exec(ctx)
	item.FileInfoId = item.FileInfo.ID.UUID
	_, err = r.db.NewInsert().Model(item.Human).Exec(ctx)
	item.HumanId = item.Human.ID
	_, err = r.db.NewInsert().Model(item).Exec(ctx)
	fmt.Println(err)
	return err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.Doctor, err error) {
	err = r.db.NewSelect().Model(&items).
		Relation("Human").
		Relation("Division").
		Relation("FileInfo").
		Order("human.surname").
		Scan(ctx)
	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.Doctor, err error) {
	err = r.db.NewSelect().Model(&item).Where("doctor.id = ?", id).
		Relation("Human").
		Relation("FileInfo").
		Relation("Division").
		Relation("DoctorComments.Comment.User").
		Scan(ctx)
	return item, err
}

func (r *Repository) getByDivisionId(ctx *gin.Context, id string) (items []models.Doctor, err error) {
	err = r.db.NewSelect().Model(&items).Where("division.id = ?", id).
		Relation("Human").
		Scan(ctx)
	return items, err
}

func (r *Repository) updateStatus(ctx *gin.Context, item *models.Doctor) (err error) {
	_, err = r.db.NewUpdate().Model(item).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Doctor{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *Repository) update(ctx *gin.Context, item *models.Doctor) (err error) {
	_, err = r.db.NewInsert().Model(item.FileInfo).
		On("CONFLICT (id) DO UPDATE").
		Set("original_name = EXCLUDED.original_name").
		Set("file_system_path = EXCLUDED.file_system_path").
		Exec(ctx)

	item.FileInfoId = item.FileInfo.ID.UUID
	r.db.NewUpdate().Model(item.Human).Where("id = ?", item.Human.ID).Exec(ctx)
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(ctx)
	return err
}

func (r *Repository) createComment(ctx *gin.Context, item *models.DoctorComment) error {
	_, err := r.db.NewInsert().Model(item.Comment).Exec(ctx)
	item.CommentId = item.Comment.ID
_, err = r.db.NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *Repository) updateComment(ctx *gin.Context, item *models.DoctorComment) error {
	_, err := r.db.NewUpdate().Model(item.Comment).Where("id = ?", item.Comment.ID).Exec(ctx)
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(ctx)
	return err
}

func (r *Repository) removeComment(ctx *gin.Context, id string) error {
	_, err := r.db.NewDelete().Model(&models.DoctorComment{}).Where("id = ?", id).Exec(ctx)
	return err
}		
