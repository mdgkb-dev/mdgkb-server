package chats

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(item *models.Chat) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (item models.ChatsWithCount, err error) {
	item.Chats = make(models.Chats, 0)
	query := r.DB().NewSelect().Model(&item.Chats).
		Relation("ChatMessages")

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(slug string) (*models.Chat, error) {
	item := models.Chat{}
	err := r.DB().NewSelect().Model(&item).Where("?TableAlias.id = ?", slug).
		Relation("ChatMessages.User.Human").
		Scan(r.ctx)
	return &item, err
}
func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.Chat{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.Chat) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
