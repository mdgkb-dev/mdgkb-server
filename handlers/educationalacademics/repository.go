package educationalacademics

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) GetAll(c context.Context) (models.EducationalAcademics, error) {
	items := make(models.EducationalAcademics, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Relation("Employee.Human.PhotoMini").
		Relation("Employee.Human.Contact.Emails").
		Relation("Employee.Human.Contact.Phones")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	err := query.Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.EducationalAcademic, error) {
	item := models.EducationalAcademic{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Doctor.Employee.Human").
		Where("?TableAlias.id = ?", id).Scan(c)
	return &item, err
}

func (r *Repository) Create(c context.Context, item *models.EducationalAcademic) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.EducationalAcademic{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.EducationalAcademic) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) UpdateAll(c context.Context, items models.EducationalAcademics) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("item_order = EXCLUDED.item_order").
		Exec(c)
	return err
}

func (r *Repository) Upsert(c context.Context, item *models.EducationalAcademic) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("item_order = EXCLUDED.item_order").
		Model(item).
		Exec(c)
	return err
}

func (r *Repository) DeleteMany(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.EducationalAcademic)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(c)
	return err
}
