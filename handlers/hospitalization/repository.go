package hospitalization

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getAll() (models.Hospitalizations, error) {
	items := make(models.Hospitalizations, 0)
	err := r.db().NewSelect().Model(&items).Relation("HospitalizationsToDocumentTypes.DocumentType").Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Hospitalization, error) {
	item := models.Hospitalization{}
	err := r.db().NewSelect().Model(&item).Where("id = ?", id).Relation("HospitalizationsToDocumentTypes.DocumentType").Scan(r.ctx)
	return &item, err
}
