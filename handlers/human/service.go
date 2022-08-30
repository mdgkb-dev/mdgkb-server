package human

import (
	"mdgkb/mdgkb-server/handlers/contactinfo"
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := contactinfo.CreateService(s.helper).Create(item.ContactInfo)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.GetFullName(), false)
	return s.repository.create(item)
}

func (s *Service) CreateMany(items models.Humans) error {
	if len(items) == 0 {
		return nil
	}
	err := contactinfo.CreateService(s.helper).CreateMany(items.GetContactInfos())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	for i := range items {
		items[i].Slug = s.helper.Util.MakeSlug(items[i].GetFullName(), false)
	}
	return s.repository.createMany(items)
}

func (s *Service) Update(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := fileinfos.CreateService(s.helper).Upsert(item.Photo)
	if err != nil {
		return err
	}
	err = fileinfos.CreateService(s.helper).Upsert(item.PhotoMini)
	if err != nil {
		return err
	}
	err = contactinfo.CreateService(s.helper).Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.GetFullName(), false)
	return s.repository.update(item)
}

func (s *Service) UpsertMany(items models.Humans) error {
	if len(items) == 0 {
		return nil
	}
	err := contactinfo.CreateService(s.helper).UpsertMany(items.GetContactInfos())
	if err != nil {
		return err
	}
	err = fileinfos.CreateService(s.helper).UpsertMany(items.GetFileInfos())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	for i := range items {
		items[i].Slug = s.helper.Util.MakeSlug(items[i].GetFullName(), true)
	}
	return s.repository.upsertMany(items)
}

func (s *Service) Upsert(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := contactinfo.CreateService(s.helper).Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	err = fileinfos.CreateService(s.helper).Upsert(item.Photo)
	if err != nil {
		return err
	}
	err = fileinfos.CreateService(s.helper).Upsert(item.PhotoMini)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	item.Slug = s.helper.Util.MakeSlug(item.GetFullName(), true)
	return s.repository.upsert(item)
}
