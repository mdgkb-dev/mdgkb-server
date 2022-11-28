package supportmessages

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

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

func (r *Repository) create(item *models.SupportMessage) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (item models.SupportMessagesWithCount, err error) {
	item.SupportMessages = make(models.SupportMessages, 0)
	query := r.db().NewSelect().
		Model(&item.SupportMessages).
		Relation("User.Human")
	r.queryFilter.HandleQuery(query)

	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id string) (*models.SupportMessage, error) {
	item := models.SupportMessage{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("User.Human").
		Where("support_messages.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.SupportMessage{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.SupportMessage) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) changeNewStatus(id string, isNew bool) (err error) {
	_, err = r.db().NewUpdate().Model(&models.SupportMessage{}).
		Set("is_new = ?", isNew).
		Where("id = ?", id).
		Exec(r.ctx)
	return err
}
