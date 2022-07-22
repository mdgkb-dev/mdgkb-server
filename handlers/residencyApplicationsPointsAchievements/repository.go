package residencyApplicationsPointsAchievements

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.ResidencyApplicationPointsAchievements) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.ResidencyApplicationPointsAchievement)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.ResidencyApplicationPointsAchievements) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("residency_application_id = EXCLUDED.residency_application_id").
		Set("points_achievement_id = EXCLUDED.points_achievement_id").
		Set("file_info_id = EXCLUDED.file_info_id").
		Set("approved = EXCLUDED.approved").
		Exec(r.ctx)
	return err
}
