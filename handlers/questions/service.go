package questions

import (
	"context"
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(c context.Context, item *models.Question) error {
	usersService := users.CreateService(s.helper)
	err := usersService.UpsertEmail(item.User)
	if err != nil {
		return err
	}

	item.User, err = usersService.Get(item.User.ID.UUID.String())
	if err != nil {
		return err
	}
	err = fileinfos.CreateService(s.helper).Create(item.File)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	err = meta.CreateService(s.helper).SendApplicationsCounts()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.QuestionsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Question, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.Question) error {
	emailStruct := struct {
		Question *models.Question
		Host     string
	}{
		item,
		s.helper.HTTP.Host,
	}
	mail, err := s.helper.Templater.ParseTemplate(emailStruct, "email/questionAnswer.gohtml")
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

func (s *Service) ReadAnswers(c context.Context, userID string) error {
	return R.ReadAnswers(c, userID)
}

func (s *Service) Publish(c context.Context, id string) error {
	return R.Publish(c, id)
}

func (s *Service) UpsertMany(c context.Context, items models.Questions) error {
	if len(items) == 0 {
		return nil
	}
	err := R.UpsertMany(c, items)
	if err != nil {
		return err
	}
	return nil
}
