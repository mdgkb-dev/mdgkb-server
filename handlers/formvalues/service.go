package formvalues

import (
	"mdgkb/mdgkb-server/handlers/children"
	"mdgkb/mdgkb-server/handlers/fieldsvalues"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Upsert(item *models.FormValue) error {
	usersService := users.CreateService(s.helper)
	err := usersService.UpsertEmail(item.User)
	if err != nil {
		return err
	}

	item.User, err = usersService.Get(item.User.ID.UUID.String())
	if err != nil {
		return err
	}
	err = children.CreateService(s.helper).Upsert(item.Child)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = fieldsvalues.CreateService(s.helper).UpsertMany(item.FieldValues)
	if err != nil {
		return err
	}
	if item.User.RejectEmail {
		return nil
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

func (s *Service) UpsertMany(items models.FormValues) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}