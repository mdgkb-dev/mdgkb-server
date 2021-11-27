package human

import (
	"mdgkb/mdgkb-server/handlers/contactInfo"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := contactInfo.CreateService(s.repository.getDB()).Create(item.ContactInfo)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	item.Slug = s.helper.MakeSlug(item.GetFullName())
	return s.repository.create(item)
}

func (s *Service) Update(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := contactInfo.CreateService(s.repository.getDB()).Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	item.Slug = s.helper.MakeSlug(item.GetFullName())
	return s.repository.update(item)
}
