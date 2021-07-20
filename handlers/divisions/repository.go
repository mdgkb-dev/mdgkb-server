package divisions

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type Repository interface {
	create(*gin.Context, *models.Division) error
	getAll(*gin.Context) ([]models.Division, error)
	updateStatus(*gin.Context, *models.Division) error
	delete(*gin.Context, string) error
}

type BRepository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *BRepository {
	return &BRepository{db}
}

func (r *BRepository) create(ctx *gin.Context, building *models.Division) (err error) {
	_, err = r.db.NewInsert().Model(building).Exec(ctx)
	return err
}

func (r *BRepository) getAll(ctx *gin.Context) ([]models.Division, error) {
	buildings := make([]models.Division, 0)

	err := r.db.NewSelect().Model(&buildings).
		Relation("Floors").
		//Relation("Floors.Divisions").
		Scan(ctx)
	return buildings, err
}

func (r *BRepository) updateStatus(ctx *gin.Context, building *models.Division) (err error) {
	_, err = r.db.NewUpdate().Model(building).Exec(ctx)
	return err
}

func (r *BRepository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Division{}).Where("id = ?", id).Exec(ctx)
	return err
}
