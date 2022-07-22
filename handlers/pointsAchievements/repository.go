package pointsAchievements

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.PointsAchievements, error) {
	items := make(models.PointsAchievements, 0)
	query := r.db().NewSelect().
		Model(&items).
		Order("points_achievements_order")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.PointsAchievement, error) {
	item := models.PointsAchievement{}
	err := r.db().NewSelect().Model(&item).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.PointsAchievement) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.PointsAchievement{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PointsAchievement) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PointsAchievements) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("points = EXCLUDED.points").
		Exec(r.ctx)
	return err
}
