package human

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Human) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAllBySlug(slug string) (models.Humans, error) {
	items := make(models.Humans, 0)
	err := r.db().NewSelect().Model(&items).Where("slug like ?", slug+"%").
		Order("slug desc").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) createMany(items models.Humans) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Human) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("humans.id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Humans) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("surname = EXCLUDED.surname").
		Set("patronymic = EXCLUDED.patronymic").
		Set("snils = EXCLUDED.snils").
		Set("photo_id = EXCLUDED.photo_id").
		Set("is_male = EXCLUDED.is_male").
		Set("date_birth = EXCLUDED.date_birth").
		Set("slug = EXCLUDED.slug").
		Set("citizenship = EXCLUDED.citizenship").
		Set("place_birth = EXCLUDED.place_birth").
		Set("car_number = EXCLUDED.car_number").
		Set("car_model = EXCLUDED.car_model").
		Set("post_index = EXCLUDED.post_index").
		Set("address = EXCLUDED.address").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Human) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("surname = EXCLUDED.surname").
		Set("patronymic = EXCLUDED.patronymic").
		Set("snils = EXCLUDED.snils").
		Set("is_male = EXCLUDED.is_male").
		Set("photo_id = EXCLUDED.photo_id").
		Set("date_birth = EXCLUDED.date_birth").
		Set("slug = EXCLUDED.slug").
		Set("citizenship = EXCLUDED.citizenship").
		Set("place_birth = EXCLUDED.place_birth").
		Set("post_index = EXCLUDED.post_index").
		Set("address = EXCLUDED.address").
		Model(item).
		Exec(r.ctx)
	return err
}
