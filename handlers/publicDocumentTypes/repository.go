package publicDocumentTypes

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.PublicDocumentType) (err error) {
	fmt.Println("item", item.Name)
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.PublicDocumentTypes, error) {
	items := make(models.PublicDocumentTypes, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("DocumentTypes.Documents.DocumentsScans.Scan").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.PublicDocumentType, error) {
	item := models.PublicDocumentType{}
	err := r.db.NewSelect().Model(&item).Where("id = ?", id).
		Relation("DocumentTypes.Documents.DocumentsScans.Scan").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.PublicDocumentType{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PublicDocumentType) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
