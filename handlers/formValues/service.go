package formValues

import (
	"mdgkb/mdgkb-server/handlers/fieldsValues"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Upsert(item *models.FormValue) error {
	err := users.CreateService(s.repository.getDB(), s.helper).UpsertEmail(item.User)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.upsert(item)
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
