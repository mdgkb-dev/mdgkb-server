package buildings

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Repository interface {
	create(*gin.Context, *models.Building) error
	getAll(*gin.Context) ([]models.Building, error)
	getByFloorId(*gin.Context, string) (models.Building, error)
	getById(*gin.Context, string) (models.Building, error)
	updateStatus(*gin.Context, *models.Building) error
	delete(*gin.Context, string) error
	update(*gin.Context, *models.Building) error
}

type BRepository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *BRepository {
	return &BRepository{db}
}

func (r *BRepository) create(ctx *gin.Context, building *models.Building) (err error) {
	_, err = r.db.NewInsert().Model(building).Exec(ctx)
	return err
}

func (r *BRepository) getAll(ctx *gin.Context) (buildings []models.Building, err error) {
	err = r.db.NewSelect().Model(&buildings).
		Relation("Floors").
		Relation("Floors.Divisions").
		Order("name").
		Order("number").
		Scan(ctx)
	return buildings, err
}

func (r *BRepository) getByFloorId(ctx *gin.Context, floorId string) (item models.Building, err error) {
	err = r.db.NewSelect().Model(&item).
		Relation("Floors").
		Relation("Floors.Divisions").
		Where("exists (select * from floors where floors.id = ? and floors.building_id = building.id)", floorId).Scan(ctx)
	return item, err
}

func (r *BRepository) getById(ctx *gin.Context, id string) (item models.Building, err error) {
	err = r.db.NewSelect().Model(&item).
		Relation("Floors").
		Relation("Floors.Divisions").
		Where("id = ?", id).Scan(ctx)
	return item, err
}

func (r *BRepository) updateStatus(ctx *gin.Context, building *models.Building) (err error) {
	_, err = r.db.NewUpdate().Model(building).Exec(ctx)
	return err
}

func (r *BRepository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Building{}).Where("id = ?", id).Exec(ctx)
	return err
}
// ! TODO Стас посмотри, плз
func (r *BRepository) update(ctx *gin.Context, building *models.Building) (err error) {
	_, err = r.db.NewUpdate().Model(building).Where("id = ?", building.ID).Exec(ctx)
	floor := new([]models.Floor)
	r.db.NewSelect().Model(floor).Where("building_id = ?", building.ID).Scan(ctx)
	for j := 0; j < len(*floor); j++ {
		found := false
		for i := 0; i < len(building.Floors); i++ {
			if building.Floors[i].ID == (*floor)[j].ID {
				found = true
			}
		}
		if !found {
			_, err = r.db.NewDelete().Model(floor).Where("id = ?", (*floor)[j].ID).Exec(ctx)
		}
	}
	for _, floors := range building.Floors {
		_, err = r.db.NewUpdate().Model(floors).Where("id = ?", floors.ID).Exec(ctx)
		if floors.ID == uuid.Nil {
			_, err = r.db.NewInsert().Model(floors).Exec(ctx)
		}
	}
	return err
}
