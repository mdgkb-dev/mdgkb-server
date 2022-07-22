package educationalOrganizationProperties

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getAll() (models.EducationalOrganizationProperties, error) {
	items := make(models.EducationalOrganizationProperties, 0)
	err := r.db().NewSelect().
		Model(&items).
		Order("educational_organization_property_order asc").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.EducationalOrganizationProperty)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationalOrganizationProperties) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("name = EXCLUDED.name").
		Set("value = EXCLUDED.value").
		Set("educational_organization_property_order = EXCLUDED.educational_organization_property_order").
		Model(&items).
		Exec(r.ctx)
	return err
}
