package valueTypes

import "mdgkb/mdgkb-server/models"

func (r *Repository) getAll() (models.ValueTypes, error) {
	items := make(models.ValueTypes, 0)
	err := r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}
