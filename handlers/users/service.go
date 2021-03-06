package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/handlers/children"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/handlers/roles"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.User) error {
	err := item.GenerateHashPassword()
	if err != nil {
		return err
	}
	err = human.CreateService(s.repository.getDB(), s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.IsActive = true
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	childrenService := children.CreateService(s.repository.getDB(), s.helper)
	err = childrenService.CreateMany(item.Children)
	if err != nil {
		return err
	}
	err = childrenService.DeleteMany(item.ChildrenForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.User) error {
	err := human.CreateService(s.repository.getDB(), s.helper).Upsert(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	//item.UUID.UUID, err = uuid.NewUUID()
	//if err != nil {
	//	return err
	//}
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	childrenService := children.CreateService(s.repository.getDB(), s.helper)
	err = childrenService.UpsertMany(item.Children)
	if err != nil {
		return err
	}
	err = childrenService.DeleteMany(item.ChildrenForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Upsert(item *models.User) error {
	err := human.CreateService(s.repository.getDB(), s.helper).Upsert(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	//item.UUID.UUID, err = uuid.NewUUID()
	//if err != nil {
	//	return err
	//}
	err = s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	childrenService := children.CreateService(s.repository.getDB(), s.helper)
	err = childrenService.UpsertMany(item.Children)
	if err != nil {
		return err
	}
	err = childrenService.DeleteMany(item.ChildrenForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertEmail(item *models.User) error {
	err := human.CreateService(s.repository.getDB(), s.helper).Upsert(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.upsertEmail(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.UsersWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.User, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetByEmail(email string) (*models.User, error) {
	item, err := s.repository.getByEmail(email)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) EmailExists(email string) (bool, error) {
	item, err := s.repository.emailExists(email)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) AddToUser(values map[string]interface{}, table string) error {
	err := s.repository.addToUser(values, table)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveFromUser(values map[string]interface{}, table string) error {
	err := s.repository.removeFromUser(values, table)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DropUUID(item *models.User) error {
	return s.repository.dropUUID(item)
}

func (s *Service) UpdatePassword(item *models.User) error {
	return s.repository.updatePassword(item)
}

func (s *Service) SetAccessLink(item *models.User) error {
	findedUser, err := s.repository.getByEmail(item.Email)
	if err != nil {
		return err
	}
	if findedUser.IsActive {
		return nil
	}
	role, err := roles.CreateService(s.repository.getDB(), s.helper).GetDefaultRole()
	if err != nil {
		return err
	}
	item.Role = role
	item.RoleID = role.ID
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	link := fmt.Sprintf("%s/access-profile/%s/%s", s.helper.HTTP.Host, item.ID, item.UUID)
	mail, err := s.helper.Templater.ParseTemplate(link, "email/profile_access.gohtml")
	if err != nil {
		return err
	}
	err = s.helper.Email.SendEmail([]string{item.Email}, "?????????????????????????? ?????????????? ?? ?????????????? ????????????", mail)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
