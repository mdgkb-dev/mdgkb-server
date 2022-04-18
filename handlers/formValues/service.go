package formValues

import (
	"mdgkb/mdgkb-server/handlers/fieldsValues"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Upsert(item *models.FormValue) error {
	usersService := users.CreateService(s.repository.getDB(), s.helper)
	err := usersService.UpsertEmail(item.User)
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
	if item.FormStatus.SendEmail {
		emailStruct := struct {
			FormValue *models.FormValue
			Host      string
		}{
			item,
			s.helper.HTTP.Host,
		}
		mail, err := s.helper.Templater.ParseTemplate(emailStruct, "email/application.gohtml")
		if err != nil {
			return err
		}
		err = s.helper.Email.SendEmail([]string{item.User.Email}, "Статус вашей заявки изменён", mail)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Get(id *string) (*models.FormValue, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
