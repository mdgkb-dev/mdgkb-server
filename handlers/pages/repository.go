package pages

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Page) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Pages, error) {
	items := make(models.Pages, 0)
	err := r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Page, error) {
	item := models.Page{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("PageDocuments.Document.FileInfo").
		Relation("PageComments.Comment").
		Where("id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Page{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Page) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getBySlug(slug *string) (*models.Page, error) {
	item := models.Page{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("PageDocuments.Document.FileInfo").
		Relation("PageComments.Comment").
		Where("slug = ?", *slug).
		Scan(r.ctx)
	return &item, err
}
