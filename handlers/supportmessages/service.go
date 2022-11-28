package supportmessages

import (
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.SupportMessage) error {
	usersService := users.CreateService(s.helper)
	err := usersService.UpsertEmail(item.User)
	if err != nil {
		return err
	}

	item.User, err = usersService.Get(item.User.ID.UUID.String())
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	err = meta.CreateService(s.helper).SendApplicationsCounts()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.SupportMessagesWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.SupportMessage, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.SupportMessage) error {
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

	return s.repository.update(item)
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) ChangeNewStatus(id string, isNew bool) error {
	return s.repository.changeNewStatus(id, isNew)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
