package models

import (
	"github.com/pro-assistance/pro-assister/tokenHelper"
)

type TokensWithUser struct {
	Tokens *tokenHelper.TokenDetails `json:"tokens"`
	User   User                      `json:"user"`
}
