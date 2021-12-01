package models

import (
	"mdgkb/mdgkb-server/helpers/tokenHelper"
)

type TokensWithUser struct {
	Token *tokenHelper.TokenDetails `json:"token"`
	User  User                      `json:"user"`
}
