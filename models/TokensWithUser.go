package models

import (
	"github.com/pro-assistance/pro-assister/tokenHelper"
)

type TokensWithUser struct {
	Tokens *tokenHelper.TokenDetails `json:"tokens"`
	User   User                      `json:"user"`
}

func (i *TokensWithUser) Init(tokens *tokenHelper.TokenDetails, user User) {
	i.Tokens = tokens
	i.User = user
}
