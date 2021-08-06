package sideOrganizations

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.SideOrganization) error
	getAll(*gin.Context) ([]models.SideOrganization, error)
	get(*gin.Context, string) (models.SideOrganization, error)
	updateStatus(*gin.Context, *models.SideOrganization) error
	delete(*gin.Context, string) error
	update(*gin.Context, *models.SideOrganization) error
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, organization *models.SideOrganization) (err error) {
	var contactInfo models.ContactInfo
	_, err = r.db.NewInsert().Model(&contactInfo).Exec(ctx)

	if err != nil {
		return err
	}

	organization.ContactInfoId = contactInfo.ID

	if organization.ContactInfo.Emails != nil {
		for _, mail := range organization.ContactInfo.Emails {
			mail.ContactInfoId = contactInfo.ID
		}

		_, err = r.db.NewInsert().Model(&organization.ContactInfo.Emails).Exec(ctx)

		if err != nil {
			return err
		}
	}

	if organization.ContactInfo.TelephoneNumbers != nil {
		for _, phone := range organization.ContactInfo.TelephoneNumbers {
			phone.ContactInfoId = contactInfo.ID
		}

		_, err = r.db.NewInsert().Model(&organization.ContactInfo.TelephoneNumbers).Exec(ctx)

		if err != nil {
			return err
		}
	}

	if organization.ContactInfo.Websites != nil {
		for _, site := range organization.ContactInfo.Websites {
			site.ContactInfoId = contactInfo.ID
		}

		_, err = r.db.NewInsert().Model(&organization.ContactInfo.Websites).Exec(ctx)

		if err != nil {
			return err
		}
	}

	_, err = r.db.NewInsert().Model(organization).Exec(ctx)
	return err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.SideOrganization, err error) {
	err = r.db.NewSelect().Model(&items).
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Scan(ctx)

	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.SideOrganization, err error) {
	err = r.db.NewSelect().Model(&item).
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Where("side_organization.id = ?", id).
		Scan(ctx)
	return item, err
}

func (r *Repository) update(ctx *gin.Context, organization *models.SideOrganization) (err error) {
	if organization.ContactInfo.Emails != nil {
		_, err = r.db.NewInsert().Model(&organization.ContactInfo.Emails).
			On("conflict (id) do update").
			Set("address = EXCLUDED.address").
			Set("description = EXCLUDED.description").
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	if organization.ContactInfo.TelephoneNumbers != nil {
		_, err = r.db.NewInsert().Model(&organization.ContactInfo.TelephoneNumbers).
			On("conflict (id) do update").
			Set("number = EXCLUDED.number").
			Set("description = EXCLUDED.description").
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	if organization.ContactInfo.Websites != nil {
		_, err = r.db.NewInsert().Model(&organization.ContactInfo.Websites).
			On("conflict (id) do update").
			Set("address = EXCLUDED.address").
			Set("description = EXCLUDED.description").
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	_, err = r.db.NewUpdate().Model(organization).Where("id = ?", organization.ID).Exec(ctx)
	return err
}

func (r *Repository) updateStatus(ctx *gin.Context, item *models.SideOrganization) (err error) {
	_, err = r.db.NewUpdate().Model(item).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	var organization models.SideOrganization

	err = r.db.NewSelect().Model(&organization).
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Where("side_organization.id = ?", id).
		Scan(ctx)

	if err != nil {
		return err
	}

	_, err = r.db.NewDelete().Model(&models.Email{}).Where("contact_info_id = ?", organization.ContactInfo.ID).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db.NewDelete().Model(&models.TelephoneNumber{}).Where("contact_info_id = ?", organization.ContactInfo.ID).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db.NewDelete().Model(&models.Website{}).Where("contact_info_id = ?", organization.ContactInfo.ID).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db.NewDelete().Model(&models.SideOrganization{}).Where("id = ?", id).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db.NewDelete().Model(&models.ContactInfo{}).Where("id = ?", organization.ContactInfo.ID).Exec(ctx)

	return err
}
