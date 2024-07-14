package projects

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) Create(c context.Context, item *models.Project) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (models.Projects, error) {
	items := make(models.Projects, 0)
	err := r.helper.DB.IDB(c).NewSelect().Model(&items).Relation("ProjectItems").Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Project, error) {
	item := models.Project{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("ProjectItems").
		Where("id = ?", id).Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Project{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Project) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) GetBySlug(c context.Context, slug string) (*models.Project, error) {
	item := models.Project{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("ProjectItems").
		Where("slug = ?", slug).
		Scan(c)
	return &item, err
}
