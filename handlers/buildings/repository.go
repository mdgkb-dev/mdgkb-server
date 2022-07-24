package buildings

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	//_ "github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(ctx *gin.Context, building *models.Building) (err error) {
	_, err = r.db().NewInsert().Model(building).Exec(ctx)
	return err
}

func (r *Repository) getAll(ctx *gin.Context) (buildings []models.Building, err error) {
	err = r.db().NewSelect().Model(&buildings).
		Relation("Floors").
		Relation("Floors.Divisions").
		Relation("Entrances").
		Relation("Entrances.Divisions").
		Order("name").
		Order("number").
		Scan(ctx)
	return buildings, err
}

func (r *Repository) getByFloorID(ctx *gin.Context, floorID string) (item models.Building, err error) {
	err = r.db().NewSelect().Model(&item).
		Relation("Floors").
		Relation("Floors.Divisions").
		Where("exists (select * from floors where floors.id = ? and floors.building_id = building.id)", floorID).Scan(ctx)
	return item, err
}

func (r *Repository) getByID(ctx *gin.Context, id string) (item models.Building, err error) {
	err = r.db().NewSelect().Model(&item).
		Relation("Floors").
		Relation("Floors.Divisions").
		Relation("Entrances").
		Relation("Entrances.Divisions").
		Where("id = ?", id).Scan(ctx)
	return item, err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Building{}).Where("id = ?", id).Exec(ctx)
	return err
}

// ! TODO Стас посмотри, плз
func (r *Repository) update(ctx *gin.Context, building *models.Building) (err error) {
	_, err = r.db().NewUpdate().Model(building).Where("id = ?", building.ID).Exec(ctx)
	if err != nil {
		return err
	}
	floor := new([]models.Floor)
	err = r.db().NewSelect().Model(floor).Where("building_id = ?", building.ID).Scan(ctx)
	if err != nil {
		return err
	}
	for j := 0; j < len(*floor); j++ {
		found := false
		for i := 0; i < len(building.Floors); i++ {
			if building.Floors[i].ID == (*floor)[j].ID {
				found = true
			}
		}
		if !found {
			_, err = r.db().NewDelete().Model(floor).Where("id = ?", (*floor)[j].ID).Exec(ctx)
			if err != nil {
				return err
			}
		}
	}
	for _, floors := range building.Floors {
		_, err = r.db().NewUpdate().Model(floors).Where("id = ?", floors.ID).Exec(ctx)
		if err != nil {
			return err
		}
		if floors.ID.UUID == uuid.Nil {
			_, err = r.db().NewInsert().Model(floors).Exec(ctx)
			if err != nil {
				return err
			}
		}
	}

	entrance := new([]models.Entrance)
	err = r.db().NewSelect().Model(entrance).Where("building_id = ?", building.ID).Scan(ctx)
	if err != nil {
		return err
	}
	for j := 0; j < len(*entrance); j++ {
		found := false
		for i := 0; i < len(building.Entrances); i++ {
			if building.Entrances[i].ID == (*entrance)[j].ID {
				found = true
			}
		}
		if !found {
			_, err = r.db().NewDelete().Model(entrance).Where("id = ?", (*entrance)[j].ID).Exec(ctx)
		}
	}
	for _, entrances := range building.Entrances {
		_, err = r.db().NewUpdate().Model(entrances).Where("id = ?", entrances.ID).Exec(ctx)
		if entrances.ID.UUID == uuid.Nil {
			_, err = r.db().NewInsert().Model(entrances).Exec(ctx)
		}
	}

	return err
}
