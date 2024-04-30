package visitsapplications

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) GetAll(c context.Context) (item models.VisitsApplicationsWithCount, err error) {
	item.VisitsApplications = make(models.VisitsApplications, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.VisitsApplications).
		Relation("Gate").
		Relation("Division").
		Relation("FormValue.Child.Human").
		Relation("FormValue.User.Human").
		Relation("Visits").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus")
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id *string) (*models.VisitsApplication, error) {
	item := models.VisitsApplication{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Gate").
		Relation("Division").
		Relation("FormValue.Child.Human").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("Visits").
		Where("visits_applications_view.id = ?", *id).Scan(c)
	return &item, err
}

func (r *Repository) Create(c context.Context, item *models.VisitsApplication) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id *string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.VisitsApplication{}).Where("id = ?", *id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.VisitsApplication) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
