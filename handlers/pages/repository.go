package pages

import (
	"context"

	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.Page) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.PagesWithCount, err error) {
	item.Pages = make(models.Pages, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.Pages)
	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id *string) (*models.Page, error) {
	item := models.Page{}
	err := r.helper.DB.IDB(c).NewSelect().
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
		Relation("Contact").
		Relation("Contact.Emails").
		Relation("Contact.PostAddresses").
		Relation("Contact.Phones").
		Relation("Contact.Websites").
		Relation("PageComments.Comment").
		Relation("Role").
		Where("id = ?", *id).Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id *string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Page{}).Where("id = ?", *id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Page) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) GetBySlug(c context.Context, slug *string) (*models.Page, error) {
	item := models.Page{}
	err := r.helper.DB.IDB(c).NewSelect().
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
		Relation("Contact").
		Relation("Contact.Emails").
		Relation("Contact.PostAddresses").
		Relation("Contact.Phones").
		Relation("Contact.Websites").
		Relation("Role").
		Where("slug = ?", *slug).
		Scan(c)
	return &item, err
}
