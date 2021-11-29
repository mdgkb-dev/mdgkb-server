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
	err = email.CreateService(s.repository.getDB()).CreateMany(item.Emails)
	if err != nil {
		return err
	}
	err = website.CreateService(s.repository.getDB()).CreateMany(item.Websites)
	if err != nil {
		return err
	}
	err = telephoneNumber.CreateService(s.repository.getDB()).CreateMany(item.TelephoneNumbers)
	if err != nil {
		return err
	}
	err = postAddress.CreateService(s.repository.getDB()).CreateMany(item.PostAddresses)
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
	err = email.CreateService(s.repository.getDB()).CreateMany(items.GetEmails())
	if err != nil {
		return err
	}
	err = website.CreateService(s.repository.getDB()).CreateMany(items.GetWebsites())
	if err != nil {
		return err
	}
	err = telephoneNumber.CreateService(s.repository.getDB()).CreateMany(items.GetTelephoneNumbers())
	if err != nil {
		return err
	}
	err = postAddress.CreateService(s.repository.getDB()).CreateMany(items.GetPostAddresses())
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
	emailService := email.CreateService(s.repository.getDB())
	err = emailService.CreateMany(item.Emails)
	if err != nil {
		return err
	}
	err = emailService.DeleteMany(item.EmailsForDelete)

	websiteService := website.CreateService(s.repository.getDB())
	err = websiteService.CreateMany(item.Websites)
	if err != nil {
		return err
	}
	err = websiteService.DeleteMany(item.WebsitesForDelete)

	telephoneNumberService := telephoneNumber.CreateService(s.repository.getDB())
	err = telephoneNumberService.CreateMany(item.TelephoneNumbers)
	if err != nil {
		return err
	}
	err = telephoneNumberService.DeleteMany(item.TelephoneNumbersForDelete)

	postAddressService := postAddress.CreateService(s.repository.getDB())
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
	emailService := email.CreateService(s.repository.getDB())
	err = emailService.UpsertMany(item.Emails)
	if err != nil {
		return err
	}
	err = emailService.DeleteMany(item.EmailsForDelete)

	websiteService := website.CreateService(s.repository.getDB())
	err = websiteService.UpsertMany(item.Websites)
	if err != nil {
		return err
	}
	err = websiteService.DeleteMany(item.WebsitesForDelete)

	telephoneNumberService := telephoneNumber.CreateService(s.repository.getDB())
	err = telephoneNumberService.UpsertMany(item.TelephoneNumbers)
	if err != nil {
		return err
	}
	err = telephoneNumberService.DeleteMany(item.TelephoneNumbersForDelete)

	postAddressService := postAddress.CreateService(s.repository.getDB())
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
	err = email.CreateService(s.repository.getDB()).UpsertMany(items.GetEmails())
	if err != nil {
		return err
	}
	err = website.CreateService(s.repository.getDB()).UpsertMany(items.GetWebsites())
	if err != nil {
		return err
	}
	err = telephoneNumber.CreateService(s.repository.getDB()).UpsertMany(items.GetTelephoneNumbers())
	if err != nil {
		return err
	}
	err = postAddress.CreateService(s.repository.getDB()).UpsertMany(items.GetPostAddresses())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
