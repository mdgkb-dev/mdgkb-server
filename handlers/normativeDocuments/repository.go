package normativeDocuments

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.NormativeDocument) error
	get(*gin.Context, string) (models.NormativeDocument, error)
	getAll(*gin.Context) ([]models.NormativeDocument, error)
	update(*gin.Context, *models.NormativeDocument) error
	delete(*gin.Context, string) error
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, document *models.NormativeDocument) (err error) {
	if document.FileInfo != nil {
		_, err = r.db.NewInsert().Model(document.FileInfo).Exec(ctx)

		if err != nil {
			return err
		}

		document.FileInfoId = document.FileInfo.ID
		_, err = r.db.NewInsert().Model(document).Column("id", "normative_document_type_id", "name", "file_info_id").Exec(ctx)
		return err
	}

	_, err = r.db.NewInsert().Model(document).Column("id", "normative_document_type_id", "name").Exec(ctx)
	return err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.NormativeDocument, err error) {
	err = r.db.NewSelect().Model(&item).
		Relation("NormativeDocumentType").
		Relation("FileInfo").
		Where("normative_document.id = ?", id).
		Scan(ctx)
	return item, err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.NormativeDocument, err error) {
	err = r.db.NewSelect().Model(&items).
		Relation("NormativeDocumentType").
		Relation("FileInfo").
		Scan(ctx)
	return items, err
}

func (r *Repository) update(ctx *gin.Context, document *models.NormativeDocument) (err error) {
	if document.FileInfo != nil {
		_, err = r.db.NewInsert().Model(document.FileInfo).
			On("conflict (id) do update").
			Set("original_name = ?", document.FileInfo.OriginalName).
			Set("file_system_path = ?", document.FileInfo.FileSystemPath).
			Exec(ctx)

		if err != nil {
			return err
		}

		document.FileInfoId = document.FileInfo.ID
	}

	_, err = r.db.NewUpdate().Model(document).Where("id = ?", document.ID).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	var document models.NormativeDocument
	err = r.db.NewSelect().Model(&document).Where("id = ?", id).Scan(ctx)

	if err != nil {
		return err
	}

	var info models.FileInfo

	if document.FileInfoId.UUID != uuid.Nil {
		err = r.db.NewSelect().Model(&info).Where("id = ?", document.FileInfoId).Scan(ctx)

		if err != nil {
			return err
		}
	}

	_, err = r.db.NewDelete().Model(&models.NormativeDocument{}).Where("id = ?", id).Exec(ctx)

	if err != nil {
		return err
	}

	if document.FileInfoId.UUID != uuid.Nil {
		_, err = r.db.NewDelete().Model(&models.FileInfo{}).Where("id = ?", info.ID).Exec(ctx)
	}

	return err
}
