package formstatusgroups

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) GetAll(c context.Context) (item models.FormStatusGroupsWithCount, err error) {
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&item.FormStatusGroups).
		Relation("FormStatuses")
	err = query.Scan(c)
	//err = query.Scan(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id *string) (*models.FormStatusGroup, error) {
	item := models.FormStatusGroup{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("FormStatuses").
		Where("form_status_groups.id = ?", *id).Scan(c)
	return &item, err
}

func (r *Repository) Upsert(c context.Context, item *models.FormStatusGroup) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(item).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("code = EXCLUDED.code").
		Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.FormStatusGroups) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("code = EXCLUDED.code").
		Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id *string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.FormStatusGroup{}).Where("id = ?", *id).Exec(c)
	return err
}
