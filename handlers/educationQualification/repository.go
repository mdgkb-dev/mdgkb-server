package educationQualification

func (r *Repository) getDB() *bun.DB {
return r.db
}

func (r *Repository) create(item *models.EducationQualification) (err error) {
_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
return err
}

func (r *Repository) getAll() (models.EducationQualifications, error) {
items := make(models.EducationQualifications, 0)
err := r.db.NewSelect().Model(&items).Scan(r.ctx)
return items, err
}

func (r *Repository) get(id *string) (*models.EducationQualification, error) {
item := models.EducationQualification{}
err := r.db.NewSelect().Model(&item).Where("id = ?", *id).Scan(r.ctx)
return &item, err
}

func (r *Repository) delete(id *string) (err error) {
_, err = r.db.NewDelete().Model(&models.EducationQualification{}).Where("id = ?", *id).Exec(r.ctx)
return err
}

func (r *Repository) update(item *models.EducationQualification) (err error) {
_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
return err
}
