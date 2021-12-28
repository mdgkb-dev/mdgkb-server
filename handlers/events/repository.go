package events

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Event) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) get(id string) (*models.Event, error) {
	item := new(models.Event)
	err := r.db.NewSelect().Model(item).
		Relation("News").
		Relation("EventApplications.FieldValues.Field").
		Relation("EventApplications.User.Human.ContactInfo.Emails").
		Relation("EventApplications.User.Human.ContactInfo.TelephoneNumbers").
		Where("events.id = ?", id).Scan(r.ctx)
	return item, err
}

func (r *Repository) update(item *models.Event) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("file_infos.id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Events) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("original_name = EXCLUDED.original_name").
		Set("file_system_path = EXCLUDED.file_system_path").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Event) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(item).
		Set("id = EXCLUDED.id").
		Exec(r.ctx)
	return err
}

//func (r *Repository) deleteMany(idPool []string) (err error) {
//	_, err = r.db.NewDelete().
//		Model((*models.DocumentType)(nil)).
//		Where("id IN (?)", bun.In(idPool)).
//		Exec(r.ctx)
//	return err
//}

func (r *Repository) createEventApplication(item *models.EventApplication) error {
	_, err := r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}
