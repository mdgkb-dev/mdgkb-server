package auth

import (
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
	return &models.TokensWithUser{Token: ts, User: *item},nil
}

func (s *Service) Login(item *models.User) (*models.TokensWithUser, error) {
	err := item.GenerateHashPassword()
	if err != nil {
		return nil, err
	}
	findedUser, err := users.CreateService(s.repository.getDB(), s.helper).GetByEmail(item.Email)
	if err != nil {
		return nil, err
	}

	if !findedUser.CompareWithHashPassword(&item.Password) {
		return nil, err
	}


	ts, err := s.helper.Token.CreateToken(findedUser.ID.String())
	if err != nil {
		return nil, err
	}

	//saveErr := helpers.CreateAuth(item.ID.String(), ts, s.redis)
	//if saveErr != nil {
	//	return nil, err
	//}

	return &models.TokensWithUser{Token: ts, User: *findedUser},nil
}