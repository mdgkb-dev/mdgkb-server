package formvalues

import (
	"context"
	"mdgkb/mdgkb-server/handlers/chats"
	"mdgkb/mdgkb-server/handlers/children"
	"mdgkb/mdgkb-server/handlers/fieldsvalues"

	// "mdgkb/mdgkb-server/handlers/formstatuses"
	"mdgkb/mdgkb-server/handlers/formvaluefiles"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) Upsert(item *models.FormValue) error {
	err := users.S.UpsertEmail(context.TODO(), item.User)
	if err != nil {
		return err
	}

	item.User, err = users.S.Get(context.TODO(), item.User.ID.UUID.String())
	if err != nil {
		return err
	}
	err = children.CreateService(s.helper).Upsert(item.Child)
	if err != nil {
		return err
	}
	var oldFormValue *models.FormValue
	if item.ID.Valid && uuid.Nil.String() != item.ID.UUID.String() {
		oldFormValue, err = s.repository.get(item.ID.UUID.String())
		if err != nil {
			return err
		}
	}
	if item.Chat == nil {
		item.Chat = &models.Chat{}
	}
	err = chats.CreateService(s.helper).Create(item.Chat)
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
	formValueFilesService := formvaluefiles.CreateService(s.helper)
	err = formValueFilesService.UpsertMany(item.FormValueFiles)
	if err != nil {
		return err
	}
	err = formValueFilesService.DeleteMany(item.FormValueFilesForDelete)
	if err != nil {
		return err
	}

	// formStatusID := item.FormStatus.ID.UUID.String()
	// formStatus, err := formstatuses.CreateService(s.helper).Get(&formStatusID)
	// if err != nil {
	// 	return err
	// }
	// если новое - слать письмо админу
	// if oldFormValue == nil {
	// 	formStatus.SentAdminEmails(item, s.helper.Email, s.helper.Templater)
	// }

	if oldFormValue != nil && oldFormValue.FormStatus != nil {
		if item.FormStatus.ID == oldFormValue.FormStatus.ID {
			return nil
		}
	}
	// if item.User.RejectEmail {
	// 	return nil
	// }
	// if item.FormStatus.SendEmail {
	// 	emailStruct := struct {
	// 		FormValue *models.FormValue
	// 		Host      string
	// 	}{
	// 		item,
	// 		s.helper.HTTP.Host,
	// 	}
	// 	mail, err := s.helper.Templater.ParseTemplate(emailStruct, "email/application.gohtml")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	err = s.helper.Email.SendEmail([]string{item.User.Email}, "Статус вашей заявки изменён", mail)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (s *Service) Get(id string) (*models.FormValue, error) {
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
