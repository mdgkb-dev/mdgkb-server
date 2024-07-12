package auth

import (
	"context"
	"fmt"
	"mdgkb/mdgkb-server/handlers/roles"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/handlers/auth"
)

func (s *Service) Register(c context.Context, item *models.User) (t *models.TokensWithUser, err error) {
	fmt.Println("register", item)
	duplicate := false
	item.UserAccountID, duplicate, err = auth.S.Register(c, item.Email, item.Password)
	item.Email = ""
	item.Password = ""
	fmt.Println("1")
	if duplicate {
		return nil, nil
	}
	fmt.Println("2")
	if err != nil {
		return nil, err
	}
	fmt.Println("3")
	role, err := roles.CreateService(s.helper).GetDefaultRole()
	if err != nil {
		return nil, err
	}
	fmt.Println("4")
	item.Role = role
	item.RoleID = role.ID
	item.IsActive = true

	item.Human = &models.Human{}
	// item.Human.ID.UUID = uuid.New()
	item.Human.ID = item.UserAccountID
	item.HumanID = item.Human.ID
	item.ID = item.UserAccountID
	fmt.Println("HumanID", item.Human.ID)
	err = users.S.Upsert(c, item)
	if err != nil {
		return nil, err
	}
	token, err := s.helper.Token.CreateToken(item)
	if err != nil {
		return nil, err
	}
	t = &models.TokensWithUser{}
	t.Init(token, *item)
	return t, err
}

func (s *Service) Login(c context.Context, email string, password string) (t *models.TokensWithUser, err error) {
	userAccountID, err, errServer := auth.S.Login(c, email, password)
	fmt.Println(userAccountID, errServer)
	if err != nil {
		return nil, err
	}
	user, err := users.S.GetByUserAccountID(c, userAccountID.UUID.String())
	if err != nil {
		return nil, err
	}
	ts, err := s.helper.Token.CreateToken(user)
	if err != nil {
		return nil, err
	}
	return &models.TokensWithUser{Tokens: ts, User: *user}, nil
}

func (s *Service) LoginAs(c context.Context, email string) (t *models.TokensWithUser, err error) {
	userAccountID, err := R.GetAccountByEmail(c, email)
	fmt.Println(userAccountID, err)
	if err != nil {
		return nil, err
	}
	user, err := users.S.GetByUserAccountID(c, userAccountID.ID.UUID.String())
	fmt.Println(user, err)
	if err != nil {
		return nil, err
	}
	ts, err := s.helper.Token.CreateToken(user)
	if err != nil {
		return nil, err
	}
	return &models.TokensWithUser{Tokens: ts, User: *user}, nil
}

func (s *Service) FindUserByEmail(c context.Context, email string) (*models.User, error) {
	findedUser, err := users.S.GetByEmail(c, email)
	if err != nil {
		return nil, err
	}
	return findedUser, nil
}

func (s *Service) GetUserByID(id string) (*models.User, error) {
	return users.S.Get(context.TODO(), id)
}

func (s *Service) DropUUID(item *models.User) error {
	return users.S.DropUUID(context.TODO(), item)
}

func (s *Service) UpdatePassword(item *models.User) error {
	return auth.S.UpdatePassword(context.TODO(), item.ID.UUID.String(), item.Password)
}

func (s *Service) RestorePassword(c context.Context, email string) error {
	userAccount, err := auth.R.GetByEmail(c, email)
	if err != nil {
		return err
	}

	emailStruct := struct {
		RestoreLink string
		Host        string
	}{
		s.helper.HTTP.GetRestorePasswordURL(userAccount.ID.UUID.String(), userAccount.UUID.String()),
		s.helper.HTTP.Host,
	}

	mail, err := s.helper.Templater.ParseTemplate(emailStruct, "email/passwordRestore.gohtml")
	if err != nil {
		return err
	}
	err = s.helper.Email.SendEmail([]string{userAccount.Email}, "Восстановление пароля", mail)
	if err != nil {
		return err
	}
	return nil
}

//
// func (s *Service) UpsertManyPathPermissions(paths models.PathPermissions) error {
// 	if len(paths) == 0 {
// 		return nil
// 	}
// 	err := s.repository.upsertManyPathPermissions(paths)
// 	if err != nil {
// 		return err
// 	}
//
// 	if len(paths.GetPathPermissionsRolesForDelete()) > 0 {
// 		err = s.repository.deleteManyPathPermissionsRoles(paths.GetPathPermissionsRolesForDelete())
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	if len(paths.GetPathPermissionsRoles()) > 0 {
// 		err = s.repository.upsertManyPathPermissionsRoles(paths.GetPathPermissionsRoles())
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }
//
// func (s *Service) GetAllPathPermissions() (models.PathPermissions, error) {
// 	return s.repository.getAllPathPermissions()
// }
//
// func (s *Service) GetAllPathPermissionsAdmin() (models.PathPermissionsWithCount, error) {
// 	return s.repository.getAllPathPermissionsAdmin()
// }
//
// func (s *Service) CheckPathPermissions(path string, roleID string) error {
// 	return s.repository.checkPathPermissions(path, roleID)
// }
//
// func (s *Service) GetPathPermissionsByRoleID(id string) (models.PathPermissions, error) {
// 	return s.repository.getPathPermissionsByRoleID(id)
// }
