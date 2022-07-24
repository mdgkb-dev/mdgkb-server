package email

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	fmt.Println(r.helper)
	fmt.Println(r.helper.DB)
	fmt.Println(*r.helper.DB.DB)
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.Emails) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Email)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Emails) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Model(&items).
		Exec(r.ctx)
	return err
}
