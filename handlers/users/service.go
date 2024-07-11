package users

import (
	"context"
	"fmt"
	"mdgkb/mdgkb-server/handlers/children"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/handlers/roles"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(c context.Context, item *models.User) error {
	err := item.GenerateHashPassword()
	if err != nil {
		return err
	}
	err = human.CreateService(s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.IsActive = true
	item.SetForeignKeys()
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	childrenService := children.CreateService(s.helper)
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

func (s *Service) Update(c context.Context, item *models.User) error {
	fmt.Println(item.Human)
	err := human.CreateService(s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	//item.UUID.UUID, err = uuid.NewUUID()
	//if err != nil {
	//	return err
	//}
	err = R.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	childrenService := children.CreateService(s.helper)
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

func (s *Service) Upsert(c context.Context, item *models.User) error {
	fmt.Println("ItemID1", item.ID, item.HumanID, item.Human)
	err := human.CreateService(s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	// item.SetForeignKeys()
	//item.UUID.UUID, err = uuid.NewUUID()
	//if err != nil {
	//	return err
	//}
	fmt.Println("ItemID2", item.ID, item.HumanID, item.Human)
	err = R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	childrenService := children.CreateService(s.helper)
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

func (s *Service) UpsertEmail(c context.Context, item *models.User) error {
	err := human.CreateService(s.helper).Upsert(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = R.UpsertEmail(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.UsersWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.User, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetByEmail(c context.Context, email string) (*models.User, error) {
	item, err := R.GetByEmail(c, email)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) EmailExists(c context.Context, email string) (bool, error) {
	item, err := R.EmailExists(c, email)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Service) AddToUser(c context.Context, values map[string]interface{}, table string) error {
	err := R.AddToUser(c, values, table)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveFromUser(c context.Context, values map[string]interface{}, table string) error {
	err := R.RemoveFromUser(c, values, table)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DropUUID(c context.Context, item *models.User) error {
	return R.DropUUID(c, item)
}

func (s *Service) UpdatePassword(c context.Context, item *models.User) error {
	return R.UpdatePassword(c, item)
}

func (s *Service) GetByUserAccountID(c context.Context, id string) (*models.User, error) {
	return R.GetByUserAccountID(c, id)
}

func (s *Service) SetAccessLink(c context.Context, item *models.User) error {
	findedUser, err := R.GetByEmail(c, item.Email)
	if err != nil {
		return err
	}
	if findedUser.IsActive {
		return nil
	}
	role, err := roles.CreateService(s.helper).GetDefaultRole()
	if err != nil {
		return err
	}
	item.Role = role
	item.RoleID = role.ID
	err = R.Update(c, item)
	if err != nil {
		return err
	}
	link := fmt.Sprintf("%s/access-profile/%s/%s", s.helper.HTTP.Host, item.ID.UUID, item.UUID)
	mail, err := s.helper.Templater.ParseTemplate(link, "email/profile_access.gohtml")
	if err != nil {
		return err
	}
	err = s.helper.Email.SendEmail([]string{item.Email}, "Подтверждение доступа к учётной записи", mail)
	if err != nil {
		return err
	}
	return nil
}
