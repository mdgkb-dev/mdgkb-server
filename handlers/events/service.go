package events

import (
	"mdgkb/mdgkb-server/handlers/fieldsValues"
	"mdgkb/mdgkb-server/handlers/forms"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Event) error {
	if item == nil {
		return nil
	}
	err := forms.CreateService(s.repository.getDB()).Upsert(item.Form)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}


func (s *Service) Get(id string) (*models.Event, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	for i := range item.EventApplications {
		item.EventApplications[i].FieldValues.PrepareValuesForPrint()
	}
	return item, err
}

func (s *Service) Update(item *models.Event) error {
	if item == nil {
		return nil
	}
	err := forms.CreateService(s.repository.getDB()).Upsert(item.Form)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return s.repository.update(item)
}

func (s *Service) UpsertMany(items models.Events) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Upsert(item *models.Event) error {
	if item == nil {
		return nil
	}
	err := forms.CreateService(s.repository.getDB()).Upsert(item.Form)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.upsert(item)
	if err != nil {
		return err
	}
	return nil
}

//func (s *Service) DeleteMany(idPool []string) error {
//	if len(idPool) == 0 {
//		return nil
//	}
//	return s.repository.deleteMany(idPool)
//}

func (s *Service) CreateEventApplication(item *models.EventApplication) error {
	err := s.repository.createEventApplication(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = fieldsValues.CreateService(s.repository.getDB()).UpsertMany(item.FieldValues)
	if err != nil {
		return err
	}
	return nil
}


