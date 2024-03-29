package auth

import (
	"errors"
	"fmt"

	"mdgkb/mdgkb-server/handlers/roles"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Register(item *models.User) (t *models.TokensWithUser, err error) {
	err = item.GenerateHashPassword()
	if err != nil {
		return nil, err
	}
	role, err := roles.CreateService(s.helper).GetDefaultRole()
	if err != nil {
		return nil, err
	}
	item.Role = role
	item.RoleID = role.ID
	item.IsActive = true
	err = users.CreateService(s.helper).Upsert(item)
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

func (s *Service) Login(item *models.Login, skipPassword bool) (t *models.TokensWithUser, err error) {
	findedUser, err := users.CreateService(s.helper).GetByEmail(item.Email)
	fmt.Println(findedUser)
	if err != nil {
		return nil, err
	}
	if !findedUser.CompareWithHashPassword(item.Password) && !skipPassword {
		return nil, errors.New("wrong login or password")
	}
	fmt.Println(findedUser)
	token, err := s.helper.Token.CreateToken(findedUser)
	if err != nil {
		return nil, err
	}
	t = &models.TokensWithUser{}
	t.Init(token, *findedUser)
	return t, err
}

func (s *Service) FindUserByEmail(email string) (*models.User, error) {
	findedUser, err := users.CreateService(s.helper).GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return findedUser, nil
}

func (s *Service) GetUserByID(id string) (*models.User, error) {
	return users.CreateService(s.helper).Get(id)
}

func (s *Service) DropUUID(item *models.User) error {
	return users.CreateService(s.helper).DropUUID(item)
}

func (s *Service) UpdatePassword(item *models.User) error {
	err := item.GenerateHashPassword()
	if err != nil {
		return err
	}
	item.IsActive = true
	return users.CreateService(s.helper).UpdatePassword(item)
}

func (s *Service) UpsertManyPathPermissions(paths models.PathPermissions) error {
	if len(paths) == 0 {
		return nil
	}
	err := s.repository.upsertManyPathPermissions(paths)
	if err != nil {
		return err
	}

	if len(paths.GetPathPermissionsRolesForDelete()) > 0 {
		err = s.repository.deleteManyPathPermissionsRoles(paths.GetPathPermissionsRolesForDelete())
		if err != nil {
			return err
		}
	}
	if len(paths.GetPathPermissionsRoles()) > 0 {
		err = s.repository.upsertManyPathPermissionsRoles(paths.GetPathPermissionsRoles())
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) GetAllPathPermissions() (models.PathPermissions, error) {
	return s.repository.getAllPathPermissions()
}

func (s *Service) GetAllPathPermissionsAdmin() (models.PathPermissionsWithCount, error) {
	return s.repository.getAllPathPermissionsAdmin()
}

func (s *Service) CheckPathPermissions(path string, roleID string) error {
	return s.repository.checkPathPermissions(path, roleID)
}

func (s *Service) GetPathPermissionsByRoleID(id string) (models.PathPermissions, error) {
	return s.repository.getPathPermissionsByRoleID(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
