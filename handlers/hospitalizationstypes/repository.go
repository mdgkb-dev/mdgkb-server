package hospitalizationstypes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.HospitalizationType) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.HospitalizationsTypes, err error) {
	query := r.db().NewSelect().Model(&items).
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.DefaultFormStatus").
		Relation("FormPattern.FormStatusGroup").
		Relation("FormPattern.Fields.ValueType").
		Relation("FormPattern.PersonalDataAgreement").
		Relation("HospitalizationTypeAnalyzes").
		Relation("HospitalizationTypeDocuments").
		Relation("HospitalizationTypeStages").
		Order("hospitalizations_types.hospitalization_type_order")
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.HospitalizationType, error) {
	item := models.HospitalizationType{}
	err := r.db().NewSelect().Model(&item).
		Where("hospitalizations_types.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.HospitalizationType{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.HospitalizationType) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
