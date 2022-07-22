package contactInfo

import (
	"mdgkb/mdgkb-server/handlers/email"
	"mdgkb/mdgkb-server/handlers/postAddress"
	"mdgkb/mdgkb-server/handlers/telephoneNumber"
	"mdgkb/mdgkb-server/handlers/website"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.ContactInfo) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = email.CreateService(s.helper).CreateMany(item.Emails)
	if err != nil {
		return err
	}
	err = website.CreateService(s.helper).CreateMany(item.Websites)
	if err != nil {
		return err
	}
	err = telephoneNumber.CreateService(s.helper).CreateMany(item.TelephoneNumbers)
	if err != nil {
		return err
	}
	err = postAddress.CreateService(s.helper).CreateMany(item.PostAddresses)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateMany(items models.ContactInfos) error {
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	err = email.CreateService(s.helper).CreateMany(items.GetEmails())
	if err != nil {
		return err
	}
	err = website.CreateService(s.helper).CreateMany(items.GetWebsites())
	if err != nil {
		return err
	}
	err = telephoneNumber.CreateService(s.helper).CreateMany(items.GetTelephoneNumbers())
	if err != nil {
		return err
	}
	err = postAddress.CreateService(s.helper).CreateMany(items.GetPostAddresses())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.ContactInfo) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	emailService := email.CreateService(s.helper)
	err = emailService.CreateMany(item.Emails)
	if err != nil {
		return err
	}
	err = emailService.DeleteMany(item.EmailsForDelete)

	websiteService := website.CreateService(s.helper)
	err = websiteService.CreateMany(item.Websites)
	if err != nil {
		return err
	}
	err = websiteService.DeleteMany(item.WebsitesForDelete)

	telephoneNumberService := telephoneNumber.CreateService(s.helper)
	err = telephoneNumberService.CreateMany(item.TelephoneNumbers)
	if err != nil {
		return err
	}
	err = telephoneNumberService.DeleteMany(item.TelephoneNumbersForDelete)

	postAddressService := postAddress.CreateService(s.helper)
	err = postAddressService.CreateMany(item.PostAddresses)
	if err != nil {
		return err
	}
	err = postAddressService.DeleteMany(item.PostAddressesForDelete)
	return nil
}

func (s *Service) Upsert(item *models.ContactInfo) error {
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	emailService := email.CreateService(s.helper)
	err = emailService.UpsertMany(item.Emails)
	if err != nil {
		return err
	}
	err = emailService.DeleteMany(item.EmailsForDelete)

	websiteService := website.CreateService(s.helper)
	err = websiteService.UpsertMany(item.Websites)
	if err != nil {
		return err
	}
	err = websiteService.DeleteMany(item.WebsitesForDelete)

	telephoneNumberService := telephoneNumber.CreateService(s.helper)
	err = telephoneNumberService.UpsertMany(item.TelephoneNumbers)
	if err != nil {
		return err
	}
	err = telephoneNumberService.DeleteMany(item.TelephoneNumbersForDelete)

	postAddressService := postAddress.CreateService(s.helper)
	err = postAddressService.UpsertMany(item.PostAddresses)
	if err != nil {
		return err
	}
	err = postAddressService.DeleteMany(item.PostAddressesForDelete)
	return nil
}

func (s *Service) UpsertMany(items models.ContactInfos) error {
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	err = email.CreateService(s.helper).UpsertMany(items.GetEmails())
	if err != nil {
		return err
	}
	err = website.CreateService(s.helper).UpsertMany(items.GetWebsites())
	if err != nil {
		return err
	}
	err = telephoneNumber.CreateService(s.helper).UpsertMany(items.GetTelephoneNumbers())
	if err != nil {
		return err
	}
	err = postAddress.CreateService(s.helper).UpsertMany(items.GetPostAddresses())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
