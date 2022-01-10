package models

import (
	"mdgkb/mdgkb-server/helpers/tokenHelper"
)

type TokensWithUser struct {
	Tokens *tokenHelper.TokenDetails `json:"tokens"`
	User   User                      `json:"user"`
}
