package projects

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Project) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Projects, error) {
	items := make(models.Projects, 0)
	err := r.db.NewSelect().Model(&items).Relation("ProjectItems").Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Project, error) {
	item := models.Project{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("ProjectItems").
		Where("id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Project{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Project) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getBySlug(slug *string) (*models.Project, error) {
	item := models.Project{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("ProjectItems").
		Where("slug = ?", *slug).
		Scan(r.ctx)
	return &item, err
}
