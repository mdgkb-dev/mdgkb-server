package supportmessages

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) Create(c context.Context, item *models.SupportMessage) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.SupportMessagesWithCount, err error) {
	item.SupportMessages = make(models.SupportMessages, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&item.SupportMessages).
		Relation("User.Human")

	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id string) (*models.SupportMessage, error) {
	item := models.SupportMessage{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("User.Human").
		Where("support_messages.id = ?", id).Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.SupportMessage{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.SupportMessage) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) ChangeNewStatus(c context.Context, id string, isNew bool) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(&models.SupportMessage{}).
		Set("is_new = ?", isNew).
		Where("id = ?", id).
		Exec(c)
	return err
}
