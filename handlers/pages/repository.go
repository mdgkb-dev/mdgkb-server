package pages

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
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

func (r *Repository) create(item *models.Page) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (item models.PagesWithCount, err error) {
	item.Pages = make(models.Pages, 0)
	query := r.db().NewSelect().Model(&item.Pages)
	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id *string) (*models.Page, error) {
	item := models.Page{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("PageSections.PageSectionDocuments.Scan").
		Relation("PageSections.PageSectionImages").
		Relation("PageSideMenus", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("page_side_menus.item_order")
		}).
		Relation("PageSideMenus.PageSections", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("page_sections.item_order")
		}).
		Relation("PageSideMenus.PageSections.PageSectionDocuments", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("page_section_documents.item_order")
		}).
		Relation("PageSideMenus.PageSections.PageSectionDocuments.Scan").
		Relation("PageSideMenus.PageSections.PageSectionImages").
		Relation("PageImages.FileInfo").
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Relation("PageComments.Comment").
		Where("id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Page{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Page) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getBySlug(slug *string) (*models.Page, error) {
	item := models.Page{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("PageSections.PageSectionDocuments.Scan").
		Relation("PageSections.PageSectionImages").
		Relation("PageSideMenus", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("page_side_menus.item_order")
		}).
		Relation("PageSideMenus.PageSections", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("page_sections.item_order")
		}).
		Relation("PageSideMenus.PageSections.PageSectionDocuments", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("page_section_documents.item_order")
		}).
		Relation("PageSideMenus.PageSections.PageSectionDocuments.Scan").
		Relation("PageSideMenus.PageSections.PageSectionImages").
		Relation("PageImages.FileInfo").
		Relation("PageComments.Comment").
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Where("slug = ?", *slug).
		Scan(r.ctx)
	return &item, err
}
