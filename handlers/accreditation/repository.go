package accreditation

func (r *Repository) getDB() *bun.DB {
return r.db
}

func (r *Repository) create(item *models.Accreditation) (err error) {
_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
return err
}

func (r *Repository) getAll() (models.Accreditations, error) {
items := make(models.Accreditations, 0)
err := r.db.NewSelect().Model(&items).Scan(r.ctx)
return items, err
}

func (r *Repository) get(id *string) (*models.Accreditation, error) {
item := models.Accreditation{}
err := r.db.NewSelect().Model(&item).Where("id = ?", *id).Scan(r.ctx)
return &item, err
}

func (r *Repository) delete(id *string) (err error) {
_, err = r.db.NewDelete().Model(&models.Accreditation{}).Where("id = ?", *id).Exec(r.ctx)
return err
}

func (r *Repository) update(item *models.Accreditation) (err error) {
_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
return err
}
