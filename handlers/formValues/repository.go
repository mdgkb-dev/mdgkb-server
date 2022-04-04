package formValues

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) upsert(item *models.FormValue) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(item).
		Set("id = EXCLUDED.id").
		Set("created_at = EXCLUDED.created_at").
		Set("is_new = EXCLUDED.is_new").
		Set("user_id = EXCLUDED.user_id").
		Exec(r.ctx)
	return err
}
