package pagesidemenus

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.PageSideMenu) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.PageSideMenus, error) {
	items := make(models.PageSideMenus, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("PageSections", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("document_types.document_type_order")
		}).
		Relation("PageSections.PageSectionDocuments", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("documents.document_order")
		}).
		Relation("PageSections.PageSectionDocuments.DocumentsScans.Scan").
		Relation("PageSections.PageSectionImages", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("document_types_images.document_type_image_order")
		}).
		Relation("PageSections.PageSectionImages.FileInfo").
		Order("public_document_type_order")

	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.PageSideMenu, error) {
	item := models.PageSideMenu{}
	err := r.db().NewSelect().Model(&item).Where("public_document_types.id = ?", id).
		Relation("PageSections", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("document_types.document_type_order")
		}).
		Relation("PageSections.PageSectionDocuments", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("documents.document_order")
		}).
		Relation("PageSections.PageSectionDocuments.DocumentsScans.Scan").
		Relation("EducationPublicDocumentType").
		Relation("PageSections.PageSectionImages", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("document_types_images.document_type_image_order")
		}).
		Relation("PageSections.PageSectionImages.FileInfo").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.PageSideMenu{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PageSideMenu) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PageSideMenus) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("item_order = EXCLUDED.item_order").
		Set("description = EXCLUDED.description").
		Set("name = EXCLUDED.name").
		Set("page_id = EXCLUDED.page_id").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.PageSideMenu)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
