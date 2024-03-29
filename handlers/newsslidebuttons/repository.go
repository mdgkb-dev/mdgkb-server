package newsslidebuttons

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.NewsSlideButtons) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.NewsSlideButton)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.NewsSlideButtons) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("name = EXCLUDED.name").
		Set("link = EXCLUDED.link").
		Set("color = EXCLUDED.color").
		Set("news_slide_button_order = EXCLUDED.news_slide_button_order").
		Set("background_color = EXCLUDED.background_color").
		Set("news_slide_id = EXCLUDED.news_slide_id").
		Set("border_color = EXCLUDED.border_color").
		Set("shadow = EXCLUDED.shadow").
		Exec(r.ctx)
	return err
}
