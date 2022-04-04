package human

import (
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Human) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) createMany(items models.Humans) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Human) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("humans.id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Humans) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("surname = EXCLUDED.surname").
		Set("patronymic = EXCLUDED.patronymic").
		Set("photo_id = EXCLUDED.photo_id").
		Set("is_male = EXCLUDED.is_male").
		Set("date_birth = EXCLUDED.date_birth").
		Set("slug = EXCLUDED.slug").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Human) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("surname = EXCLUDED.surname").
		Set("patronymic = EXCLUDED.patronymic").
		Set("is_male = EXCLUDED.is_male").
		Set("photo_id = EXCLUDED.photo_id").
		Set("date_birth = EXCLUDED.date_birth").
		Set("slug = EXCLUDED.slug").
		Model(item).
		Exec(r.ctx)
	return err
}
