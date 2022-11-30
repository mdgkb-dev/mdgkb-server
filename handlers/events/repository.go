package events

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Event) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) get(id string) (*models.Event, error) {
	item := new(models.Event)
	err := r.db().NewSelect().Model(item).
		Relation("News").
		Relation("EventApplications.FieldValues.Field").
		Relation("EventApplications.User.Human.ContactInfo.Emails").
		Relation("EventApplications.User.Human.ContactInfo.TelephoneNumbers").
		Where("events.id = ?", id).Scan(r.ctx)
	return item, err
}

func (r *Repository) update(item *models.Event) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("file_infos.id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Events) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("start_date = EXCLUDED.start_date").
		Set("end_date = EXCLUDED.end_date").
		Set("form_id = EXCLUDED.form_id").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Event) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(item).
		Set("id = EXCLUDED.id").
		Set("start_date = EXCLUDED.start_date").
		Set("end_date = EXCLUDED.end_date").
		Set("form_id = EXCLUDED.form_id").
		Exec(r.ctx)
	return err
}

//func (r *Repository) deleteMany(idPool []string) (err error) {
//	_, err = r.db().NewDelete().
//		Model((*models.PageSection)(nil)).
//		Where("id IN (?)", bun.In(idPool)).
//		Exec(r.ctx)
//	return err
//}

func (r *Repository) createEventApplication(item *models.EventApplication) error {
	_, err := r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAllForMain() (models.Events, error) {
	items := make(models.Events, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("News").
		//Relation("Event").
		//Join("JOIN events on news.event_id = event.id").
		Order("news.published_on DESC").
		Limit(12).
		Scan(r.ctx)
	return items, err
}
