package sideorganizations

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	//_ "github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(ctx *gin.Context, organization *models.SideOrganization) (err error) {
	var contactInfo models.ContactInfo
	_, err = r.db().NewInsert().Model(&contactInfo).Exec(ctx)
	if err != nil {
		return err
	}

	organization.ContactInfoID = contactInfo.ID

	if len(organization.ContactInfo.Emails) > 0 {
		for _, mail := range organization.ContactInfo.Emails {
			mail.ContactInfoID = contactInfo.ID
		}

		_, err = r.db().NewInsert().Model(&organization.ContactInfo.Emails).Exec(ctx)

		if err != nil {
			return err
		}
	}

	if len(organization.ContactInfo.PostAddresses) > 0 {
		for _, address := range organization.ContactInfo.PostAddresses {
			address.ContactInfoID = contactInfo.ID
		}

		_, err = r.db().NewInsert().Model(&organization.ContactInfo.PostAddresses).Exec(ctx)

		if err != nil {
			return err
		}
	}

	if len(organization.ContactInfo.TelephoneNumbers) > 0 {
		for _, phone := range organization.ContactInfo.TelephoneNumbers {
			phone.ContactInfoID = contactInfo.ID
		}

		_, err = r.db().NewInsert().Model(&organization.ContactInfo.TelephoneNumbers).Exec(ctx)

		if err != nil {
			return err
		}
	}

	if len(organization.ContactInfo.Websites) > 0 {
		for _, site := range organization.ContactInfo.Websites {
			site.ContactInfoID = contactInfo.ID
		}

		_, err = r.db().NewInsert().Model(&organization.ContactInfo.Websites).Exec(ctx)

		if err != nil {
			return err
		}
	}

	_, err = r.db().NewInsert().Model(organization).Exec(ctx)
	return err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.SideOrganization, err error) {
	err = r.db().NewSelect().Model(&items).
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Order("side_organization.name").
		Scan(ctx)

	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.SideOrganization, err error) {
	err = r.db().NewSelect().Model(&item).
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Where("side_organization.id = ?", id).
		Scan(ctx)
	return item, err
}

func (r *Repository) update(ctx *gin.Context, organization *models.SideOrganization) (err error) {
	organization.ContactInfoID = organization.ContactInfo.ID
	var existingOrg models.SideOrganization

	_ = r.db().NewSelect().Model(&existingOrg).
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Where("side_organization.id = ?", organization.ID).
		Scan(ctx)

	if len(organization.ContactInfo.Emails) > 0 {
		incomingSet := map[uuid.UUID]bool{}

		for _, mail := range organization.ContactInfo.Emails {
			mail.ContactInfoID = organization.ContactInfo.ID
			incomingSet[mail.ID] = true
		}

		deleteQuery := r.db().NewDelete().Model(&models.Email{}).Where("1 = 0")

		for _, mail := range existingOrg.ContactInfo.Emails {
			if !incomingSet[mail.ID] {
				deleteQuery = deleteQuery.WhereOr("id = ?", mail.ID)
			}
		}

		_, err = deleteQuery.Exec(ctx)

		if err != nil {
			return err
		}

		_, err = r.db().NewInsert().Model(&organization.ContactInfo.Emails).
			On("conflict (id) do update").
			Set("address = EXCLUDED.address").
			Set("description = EXCLUDED.description").
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	if len(organization.ContactInfo.PostAddresses) > 0 {
		incomingSet := map[uuid.UUID]bool{}

		for _, address := range organization.ContactInfo.PostAddresses {
			address.ContactInfoID = organization.ContactInfo.ID
			incomingSet[address.ID] = true
		}

		deleteQuery := r.db().NewDelete().Model(&models.PostAddress{}).Where("1 = 0")

		for _, address := range existingOrg.ContactInfo.PostAddresses {
			if !incomingSet[address.ID] {
				deleteQuery = deleteQuery.WhereOr("id = ?", address.ID)
			}
		}

		_, err = deleteQuery.Exec(ctx)

		if err != nil {
			return err
		}

		_, err = r.db().NewInsert().Model(&organization.ContactInfo.PostAddresses).
			On("conflict (id) do update").
			Set("address = EXCLUDED.address").
			Set("description = EXCLUDED.description").
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	if len(organization.ContactInfo.TelephoneNumbers) > 0 {
		incomingSet := map[uuid.UUID]bool{}

		for _, phone := range organization.ContactInfo.TelephoneNumbers {
			phone.ContactInfoID = organization.ContactInfo.ID
			incomingSet[phone.ID] = true
		}

		deleteQuery := r.db().NewDelete().Model(&models.TelephoneNumber{}).Where("1 = 0")

		for _, phone := range existingOrg.ContactInfo.TelephoneNumbers {
			if !incomingSet[phone.ID] {
				deleteQuery = deleteQuery.WhereOr("id = ?", phone.ID)
			}
		}

		_, err = deleteQuery.Exec(ctx)

		if err != nil {
			return err
		}

		_, err = r.db().NewInsert().Model(&organization.ContactInfo.TelephoneNumbers).
			On("conflict (id) do update").
			Set("number = EXCLUDED.number").
			Set("description = EXCLUDED.description").
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	if len(organization.ContactInfo.Websites) > 0 {
		incomingSet := map[uuid.UUID]bool{}

		for _, site := range organization.ContactInfo.Websites {
			site.ContactInfoID = organization.ContactInfo.ID
			incomingSet[site.ID] = true
		}

		deleteQuery := r.db().NewDelete().Model(&models.Website{}).Where("1 = 0")

		for _, site := range existingOrg.ContactInfo.Websites {
			if !incomingSet[site.ID] {
				deleteQuery = deleteQuery.WhereOr("id = ?", site.ID)
			}
		}

		_, err = deleteQuery.Exec(ctx)

		if err != nil {
			return err
		}

		_, err = r.db().NewInsert().Model(&organization.ContactInfo.Websites).
			On("conflict (id) do update").
			Set("address = EXCLUDED.address").
			Set("description = EXCLUDED.description").
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	_, err = r.db().NewUpdate().Model(organization).Where("id = ?", organization.ID).Exec(ctx)
	return err
}

func (r *Repository) updateStatus(ctx *gin.Context, item *models.SideOrganization) (err error) {
	_, err = r.db().NewUpdate().Model(item).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	var organization models.SideOrganization

	err = r.db().NewSelect().Model(&organization).
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.PostAddresses").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactInfo.Websites").
		Where("side_organization.id = ?", id).
		Scan(ctx)

	if err != nil {
		return err
	}

	_, err = r.db().NewDelete().Model(&models.Email{}).Where("contact_info_id = ?", organization.ContactInfo.ID).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db().NewDelete().Model(&models.PostAddress{}).Where("contact_info_id = ?", organization.ContactInfo.ID).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db().NewDelete().Model(&models.TelephoneNumber{}).Where("contact_info_id = ?", organization.ContactInfo.ID).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db().NewDelete().Model(&models.Website{}).Where("contact_info_id = ?", organization.ContactInfo.ID).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db().NewDelete().Model(&models.SideOrganization{}).Where("id = ?", id).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db().NewDelete().Model(&models.ContactInfo{}).Where("id = ?", organization.ContactInfo.ID).Exec(ctx)

	return err
}
