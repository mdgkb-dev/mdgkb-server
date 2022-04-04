package human

import (
	"mdgkb/mdgkb-server/handlers/contactInfo"
	"mdgkb/mdgkb-server/handlers/fileInfos"
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

func (s *Service) CreateMany(items models.Humans) error {
	if len(items) == 0 {
		return nil
	}
	err := contactInfo.CreateService(s.repository.getDB()).CreateMany(items.GetContactInfos())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	for i := range items {
		items[i].Slug = s.helper.MakeSlug(items[i].GetFullName())
	}
	return s.repository.createMany(items)
}

func (s *Service) Update(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := fileInfos.CreateService(s.repository.getDB()).Upsert(item.Photo)
	if err != nil {
		return err
	}
	err = contactInfo.CreateService(s.repository.getDB()).Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	item.Slug = s.helper.MakeSlug(item.GetFullName())
	return s.repository.update(item)
}

func (s *Service) UpsertMany(items models.Humans) error {
	if len(items) == 0 {
		return nil
	}
	err := contactInfo.CreateService(s.repository.getDB()).UpsertMany(items.GetContactInfos())
	if err != nil {
		return err
	}
	err = fileInfos.CreateService(s.repository.getDB()).UpsertMany(items.GetPhotos())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	for i := range items {
		items[i].Slug = s.helper.MakeSlug(items[i].GetFullName())
	}
	return s.repository.upsertMany(items)
}

func (s *Service) Upsert(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := contactInfo.CreateService(s.repository.getDB()).Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	err = fileInfos.CreateService(s.repository.getDB()).Upsert(item.Photo)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	item.Slug = s.helper.MakeSlug(item.GetFullName())
	return s.repository.upsert(item)
}
