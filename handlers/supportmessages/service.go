package supportmessages

import (
	"context"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(c context.Context, item *models.SupportMessage) error {
	err := users.S.UpsertEmail(context.TODO(), item.User)
	if err != nil {
		return err
	}

	item.User, err = users.S.Get(context.TODO(), item.User.ID.UUID.String())
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	err = meta.S.SendApplicationsCounts()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.SupportMessagesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.SupportMessage, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.SupportMessage) error {
	emailStruct := struct {
		SupportMessage *models.SupportMessage
		Host           string
	}{
		item,
		s.helper.HTTP.Host,
	}
	mail, err := s.helper.Templater.ParseTemplate(emailStruct, "email/supportMessageAnswer.gohtml")
	if err != nil {
		return err
	}
	err = s.helper.Email.SendEmail([]string{item.User.Email}, "Ответ на Ваш вопрос на сайте МДГКБ", mail)
	if err != nil {
		return err
	}

	return R.Update(c, item)
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) ChangeNewStatus(c context.Context, id string, isNew bool) error {
	return R.ChangeNewStatus(c, id, isNew)
}
