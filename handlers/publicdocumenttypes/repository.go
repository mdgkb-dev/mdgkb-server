package publicdocumenttypes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.PublicDocumentType) (err error) {
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

func (r *Repository) getAll() (models.PublicDocumentTypes, error) {
	items := make(models.PublicDocumentTypes, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("DocumentTypes", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("document_types.document_type_order")
		}).
		Relation("DocumentTypes.Documents", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("documents.document_order")
		}).
		Relation("DocumentTypes.Documents.DocumentsScans.Scan").
		Relation("DocumentTypes.DocumentTypeImages", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("document_types_images.document_type_image_order")
		}).
		Relation("DocumentTypes.DocumentTypeImages.FileInfo").
		Order("public_document_type_order")

	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.PublicDocumentType, error) {
	item := models.PublicDocumentType{}
	err := r.db().NewSelect().Model(&item).Where("public_document_types.id = ?", id).
		Relation("DocumentTypes", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("document_types.document_type_order")
		}).
		Relation("DocumentTypes.Documents", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("documents.document_order")
		}).
		Relation("DocumentTypes.Documents.DocumentsScans.Scan").
		Relation("EducationPublicDocumentType").
		Relation("DocumentTypes.DocumentTypeImages", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("document_types_images.document_type_image_order")
		}).
		Relation("DocumentTypes.DocumentTypeImages.FileInfo").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.PublicDocumentType{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PublicDocumentType) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PublicDocumentTypes) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("public_document_type_order = EXCLUDED.public_document_type_order").
		Exec(r.ctx)
	return err
}
