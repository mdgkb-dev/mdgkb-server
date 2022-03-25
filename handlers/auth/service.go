package auth

import (
	"fmt"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Register(item *models.User) (*models.TokensWithUser, error) {
	err := item.GenerateHashPassword()
	if err != nil {
		return nil, err
	}
	err = users.CreateService(s.repository.getDB(), s.helper).Create(item)
	if err != nil {
		return nil, err
	}
	ts, err := s.helper.Token.CreateToken(item.ID.String())
	if err != nil {
		return nil, err
	}
	return &models.TokensWithUser{Tokens: ts, User: *item}, nil
}

func (s *Service) Login(item *models.User) (*models.TokensWithUser, error) {
	//err := item.GenerateHashPassword()
	//if err != nil {
	//	return nil, err
	//}
	findedUser, err := users.CreateService(s.repository.getDB(), s.helper).GetByEmail(item.Email)
	if err != nil {
		return nil, err
	}
	if !findedUser.CompareWithHashPassword(item.Password) {
		fmt.Println(findedUser.CompareWithHashPassword(item.Password))
		return nil, err
	}
	//fmt.Println(item.Password)
	ts, err := s.helper.Token.CreateToken(findedUser.ID.String())
	if err != nil {
		return nil, err
	}

	//saveErr := helpers.CreateAuth(item.ID.String(), ts, s.redis)
	//if saveErr != nil {
	//	return nil, err
	//}

	return &models.TokensWithUser{Tokens: ts, User: *findedUser}, nil
}

func (s *Service) FindUserByEmail(email string) (*models.User, error) {
	findedUser, err := users.CreateService(s.repository.getDB(), s.helper).GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return findedUser, nil
}

func (s *Service) GetUserByID(id string) (*models.User, error) {
	return users.CreateService(s.repository.getDB(), s.helper).Get(id)
}

func (s *Service) DropUUID(item *models.User) error {
	return users.CreateService(s.repository.getDB(), s.helper).DropUUID(item)
}

func (s *Service) UpdatePassword(item *models.User) error {
	err := item.GenerateHashPassword()
	if err != nil {
		return err
	}
	return users.CreateService(s.repository.getDB(), s.helper).UpdatePassword(item)
}
