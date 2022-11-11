package employees

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
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.Employee) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.Employees, err error) {
	query := r.db().NewSelect().Model(&items).
		Relation("Human").
		Relation("Human.PhotoMini")
	r.queryFilter.HandleQuery(query)
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(slug string) (*models.Employee, error) {
	item := models.Employee{}
	err := r.db().NewSelect().Model(&item).Where("doctors_view.slug = ?", slug).
		Relation("Human.Photo").
		Relation("Human.PhotoMini").
		Relation("Regalias").
		Relation("Experiences").
		Relation("Certificates.Scan").
		Relation("Educations.EducationCertification").
		Relation("Educations.EducationAccreditation").
		Relation("DoctorComments.Comment", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("comment.published_on DESC")
		}).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Doctor{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Employee) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
